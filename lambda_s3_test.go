package test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestTerraformLambdaS3Workflow(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to the Terraform code that will be tested.
		TerraformDir: "terraform",

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"terraform.tfvars"},
	}

	//Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply".
	terraform.InitAndApply(t, terraformOptions)

	// Set up the AWS session
	creds := credentials.NewSharedCredentials("terraform/credentials", "default")
	config := &aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: creds,
	}
	sess := session.New(config)
	svc := s3.New(sess)

	// Download the logo we want to resize.
	resp, err := http.Get("https://www.google.co.uk/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png")
	if err != nil {
		log.Print(err)
	}

	// Now upload the logo image to the src bucket, and verify the resized version is present in the dst bucket.
	logoTransformerSrcBucketName := terraform.Output(t, terraformOptions, "codescreen-logos-bucket-name")
	logoTransformerDstBucketName := terraform.Output(t, terraformOptions, "codescreen-logos-resized-bucket-name")

	params := &s3manager.UploadInput{
		Bucket: aws.String(logoTransformerSrcBucketName),
		Key:    aws.String("googleLogo.png"),
		Body:   resp.Body,
	}

	uploader := s3manager.NewUploader(sess)
	uploader.Upload(params)

	time.Sleep(10 * time.Second) // Sleep for 10 seconds to give the workflow enough time to complete.

	headObjectOutput, headObjectErr := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(logoTransformerDstBucketName),
		Key:    aws.String("resized-googleLogo.png"),
	})

	log.Print(headObjectOutput)
	assert.Equal(t, nil, headObjectErr)

	// Finally, delete the two image files so that Terraform can delete the S3 buckets as part of the destroy phase.
	deleteFileInput := &s3.DeleteObjectInput{
		Bucket: aws.String(logoTransformerSrcBucketName),
		Key:    aws.String("googleLogo.png"),
	}

	deleteFileResizedInput := &s3.DeleteObjectInput{
		Bucket: aws.String(logoTransformerDstBucketName),
		Key:    aws.String("resized-googleLogo.png"),
	}

	svc.DeleteObject(deleteFileInput)
	svc.DeleteObject(deleteFileResizedInput)
}
