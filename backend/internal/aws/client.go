// File: backend/internal/aws/client.go
package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// Credentials represents AWS access credentials
type Credentials struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

// VerifyCredentials checks if AWS credentials are valid by making a simple API call
func VerifyCredentials(creds Credentials) error {
	// Create AWS session with the provided credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(creds.Region),
		Credentials: credentials.NewStaticCredentials(creds.AccessKey, creds.SecretKey, ""),
	})

	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Use IAM service to retrieve the user (a simple API call to verify credentials)
	svc := iam.New(sess)
	_, err = svc.GetUser(&iam.GetUserInput{})

	if err != nil {
		return fmt.Errorf("invalid AWS credentials: %w", err)
	}

	return nil
}

// Additional AWS functionality can be added here as needed
// For example, functions to fetch EC2 instances, S3 buckets, etc.