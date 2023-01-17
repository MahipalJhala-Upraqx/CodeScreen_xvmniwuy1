// TODO Implement
provider "aws" {
  profile = "default"
  shared_credentials_file ="./credentials"
  region  = var.region
}

resource "aws_lambda_function" "logo_transformer" {
  filename         = "logo-transformer-package.zip"
  function_name    = "logo-transformer-${var.suffix}"
  handler          = "index.handler"
  role             = var.permissions_boundary_arn
  runtime          = "nodejs16.x"
  source_code_hash = filebase64sha256("logo-transformer-package.zip")
  timeout          = 900
}

# Creating s3 resource for invoking to lambda function
resource "aws_s3_bucket" "bucket" {
  bucket = "codescreen-logos-${var.suffix}"
  acl    = "public-read-write"
}

# Creating s3 resource for invoking to lambda function
resource "aws_s3_bucket" "destinationbucket" {
  bucket = "${aws_s3_bucket.bucket.id}-resized"
  acl    = "public-read-write"
}

# Adding S3 bucket as trigger to my lambda and giving the permissions
resource "aws_s3_bucket_notification" "aws-lambda-trigger" {
  bucket = aws_s3_bucket.bucket.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.logo_transformer.arn
    events              = ["s3:ObjectCreated:*"]

  }
}
resource "aws_lambda_permission" "test" {
  statement_id  = "AllowS3Invoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.logo_transformer.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = "arn:aws:s3:::${aws_s3_bucket.bucket.id}"
}
