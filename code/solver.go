// © 2019-2020 nextmv.io inc. All rights reserved.
// nextmv.io, inc. CONFIDENTIAL
//
// This file includes unpublished proprietary source code of nextmv.io, inc.
// The copyright notice above does not evidence any actual or intended
// publication of such source code. Disclosure of this source code or any
// related proprietary information is strictly prohibited without the express
// written permission of nextmv.io, inc.

package samstep

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nextmv-io/hop/model"
	"github.com/nextmv-io/hop/solve"
	"github.com/nextmv-io/hop/solve/diagram/restrict"
)

type S3Object struct {
	Bucket string
	Key    string
}

// Solver solves the N queens problem.
func Solver(n int, opt solve.Options) (solve.Solver, error) {
	// Download the file from s3
	// n, err := fetchS3(in)
	// if err != nil {
	// 	return nil, err
	// }

	// Random restriction makes it easier to find lots of layouts.
	opt.Diagram.Restrictor = restrict.Random()

	// In the root state, no rows have been assigned a queen.
	b := board{rows: model.Domains(n, model.Range(0, n-1))}
	return solve.Satisfier(b, opt), nil
}

func fetchS3(from S3Object) (int, error) {
	sess := session.Must(session.NewSession())

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	f := aws.NewWriteAtBuffer([]byte{})

	// Write the contents of S3 Object to the file
	_, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(from.Bucket),
		Key:    aws.String(from.Key),
	})
	if err != nil {
		return 0, fmt.Errorf("failed to download file, %v", err)
	}

	var val int
	err = json.Unmarshal(f.Bytes(), &val)
	if err != nil {
		return 0, err
	}
	return val, nil
}
