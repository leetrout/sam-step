package main

import (
	"github.com/nextmv-io/hop/run/lambda/s3"
	"stub.local/samstep"
)

func main() {
	s3.Run(samstep.Solver)
}
