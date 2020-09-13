package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
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
			os.Exit(1)
		}
		// https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-data.html
		output := fmt.Sprintf(`{"bucket": %q, "key": %q}`, s3bucket, s3key)
		fmt.Println(output)
	} else {
		fmt.Println(string(jsOut))
	}
}
