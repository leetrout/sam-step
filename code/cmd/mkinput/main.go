package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"samstep"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func mkinput() (samstep.S3Object, error) {
	output := samstep.S3Object{}
	i := 4
	jsOut, _ := json.Marshal(i)
	s3bucket := os.Getenv("S3_OUTPUT_BUCKET")
	s3key := strings.Replace(time.Now().Format(time.RFC3339), ":", "", -1)
	if s3bucket != "" {
		sess := session.Must(session.NewSession())
		uploader := s3manager.NewUploader(sess)

		// Upload the file to S3.
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(s3bucket),
			Key:    aws.String(s3key),
			Body:   bytes.NewReader(jsOut),
		})
		if err != nil {
			fmt.Printf("failed to upload file: %v\n", err)
			return output, err
		}
		// https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-data.html
		output.Bucket = s3bucket
		output.Key = s3key
		return output, nil
	}

	return output, nil
}

func main() {
	lambda.Start(mkinput)
}
