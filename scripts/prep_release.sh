#!/bin/bash

PATH=$GOPATH/bin/:$PATH
VERSION=$1
TIME=$(date +%s)

if [ -z "$VERSION" ]; then
	echo "You must call this script with a version as first argument"
	exit 1
fi
rm -rf build
mkdir build

function make_release() {
	NAME=$1
	export GOOS=$2
	export GOARCH=$3
	SUFFIX=$4
	RELEASE_PATH=build/cwput-$NAME-${VERSION}
	RELEASE_FILE=cwput-$NAME-${VERSION}.tar.gz
	mkdir $RELEASE_PATH
	cp README.md $RELEASE_PATH
	cp LICENSE.txt $RELEASE_PATH
	echo $GOOS
	echo $GOARCH
	go build -o $RELEASE_PATH/cwput${SUFFIX} *.go
	PREV_WD=$(PWD)
	cd  $RELEASE_PATH
	tar cvfz ../$RELEASE_FILE .
	cd ../../
}

make_release "osx" "darwin" "amd64" ""
make_release "linux-amd64" "linux" "amd64" ""
make_release "windows-amd64" "windows" "amd64" ".exe"