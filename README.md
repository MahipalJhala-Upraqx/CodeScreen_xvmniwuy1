# Mahipal Upraqx's assessment
[Lambda](https://aws.amazon.com/lambda/) is a serverless compute platform provided by [AWS](https://aws.amazon.com/).

[S3](https://aws.amazon.com/s3/) is an object storage service also provided by AWS.

[terraform/functions/logo-transformer](terraform/functions/logo-transformer/) is a `Node.js` utility that
resizes image logos. It reads an image file from one S3 bucket (the `source` bucket), resizes the image, and then uploads the resized image into a separate
S3 bucket (the `destination` bucket).

The goal of this test is to use [Terraform](https://www.terraform.io/) to productionize this resize logo image workflow by
deploying the
logo-transformer code into an `AWS Lambda` function, and have the Lambda function triggered every time a new
logo image is uploaded to an S3 bucket.

## Your Task

You are required to implement the [main.tf](terraform/main.tf) file so that when you run `terraform apply`, the
logo-transformer code is deployed inside a Lambda function, and the deployed lambda function is triggered every time a new
object is added to the S3 bucket.

The logo-transformer code has already been packaged up into a deployment package Zip file. It can be download from
[here](https://codescreen-assessment-packages.s3-eu-west-1.amazonaws.com/logo-transformer-package.zip).

A `IAM User` has been set up for you in our AWS account, and the credentials are located
[here](terraform/credentials). This user already has all permissions needed to successfully create all the AWS resources required
for this task.

All `Input Variables` required for this task have also already been set in the [terraform.tfvars](terraform/terraform.tfvars) file.

The [lambda_s3_test.go](lambda_s3_test.go) unit test file should pass if your solution has been implemented correctly.

## Requirements

For security reasons, any `IAM Role` you create **must** include the [Permissions boundary](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) that
we have attached to your IAM User. The `ARN` for this Permission boundary `IAM Policy` is included in the [terraform.tfvars](terraform/terraform.tfvars) file.

The Lambda function name must begin with `logo-transformer-{suffix}` and the source S3 bucket name must begin with `codescreen-logos-{suffix}`,
where `{suffix}` refers to the suffix variable in [terraform.tfvars](terraform/terraform.tfvars).

The destination S3 bucket name must be of the form `{sourceBucketName}-resized`, where `{sourceBucketName}` refers to the name of the source S3 bucket.

The [lambda_s3_test.go](lambda_s3_test.go), [go.mod](go.mod), [terraform.tfvars](terraform/terraform.tfvars)
and [variables.tf](terraform/variables.tf) files **must not** be modified.

The [logo-transformer](terraform/functions/logo-transformer/) code also **must not** be edited.

The [logo-transformer-package.zip](https://codescreen-assessment-packages.s3-eu-west-1.amazonaws.com/logo-transformer-package.zip) file **must** be commited to
this repo as part of your solution.

Your IAM user does not have console access, so we recommend using the [Lumigo CLI](https://github.com/lumigo-io/lumigo-CLI#lumigo-cli-tail-cloudwatch-logs) to tail the cloudwatch logs for your lambda
function when testing your solution.

## Tests

To run the tests, you must first install `Go` on your system. See [here](https://golang.org/doc/install).

Once `Go` is installed, run `go mod download` then `go mod tidy` to install the test dependencies, and then run `go test` to run the tests.

##

This test should take no longer than 3 hours to complete successfully.

Good luck!
## License

At CodeScreen, we strongly value the integrity and privacy of our assessments. As a result, this repository is under exclusive copyright, which means you **do not** have permission to share your solution to this test publicly (i.e., inside a public GitHub/GitLab repo, on Reddit, etc.). <br>

## Submitting your solution

Please push your changes to the `main branch` of this repository. You can push one or more commits. <br>

Once you are finished with the task, please click the `Submit Solution` link on <a href="https://app.codescreen.com/candidate/17b55331-48d6-491d-83c6-c54623a4b530" target="_blank">this screen</a>.