.PHONY: build build-go build-MakeInputFunction

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build-go:
	cd ../code/cmd/mkinput && go build -o ../../../aws-lambda/build/mkinput
	cd build && zip mkinput.zip mkinput

build-MakeInputFunction:
	cd ${ROOT_DIR}/../code/cmd/mkinput && go build -o ${ARTIFACTS_DIR}/mkinput

build: build-go
	sam build
