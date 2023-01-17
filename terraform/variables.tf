variable "suffix" {
  description = "The suffix used for S3 bucket names and Lambda function names."
  type        = string
}

variable "permissions_boundary_arn" {
  description = "The ARN of the permission boundary policy"
  type        = string
}

variable "region" {
  description = "The region in which to deploy the resources."
  type        = string
}
