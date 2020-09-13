package main

import (
	"github.com/nextmv-io/hop/run/lambda"
	"stub.local/samstep"
)

func main() {
	lambda.Run(samstep.Solver)
}
