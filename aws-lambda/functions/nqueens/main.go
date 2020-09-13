package main

import (
	"stub.local/samstep"

	"github.com/nextmv-io/hop/run/lambda/s3"
)

func main() {
	s3.Run(samstep.Solver)
}
