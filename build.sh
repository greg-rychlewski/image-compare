#!/bin/bash

# Update this variable when publishing a new version
APP_VERSION=1.0.0

# This will embed develpment information into our app
GIT_HASH=`git rev-parse HEAD`
BUILD_TIME_UTC=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
GO_VERSION=`go version`

# List all target operating systems and architectures in space separated lists
TARGET_OS='windows darwin'
TARGET_ARCH='386 amd64'

# Run unit tests before building
go test -v ./...

rc=$?

if [ $rc != 0 ]; then
  exit $rc
fi

# Build app for each operating system/architecture combination
for OS in ${TARGET_OS}; do
  for ARCH in ${TARGET_ARCH}; do
    BINARY_FILE="image_compare-${OS}-${ARCH}"

    if [[ ${OS} == 'windows' ]]; then
      BINARY_FILE=${BINARY_FILE}.exe
    fi;

    GOOS=${OS} GOARCH=${ARCH} \
    go build \
      -o bin/${BINARY_FILE}\
      -ldflags "-X 'main.goBuildVersion=${GO_VERSION}' -X main.buildTime=${BUILD_TIME_UTC} -X main.gitHash=${GIT_HASH} -X main.version=${APP_VERSION}"

    rc=$?

    if [ $rc != 0 ]; then
      exit $rc
    fi

    # Rename file to something more user-friendly and tar it so user can preserve permissions when extracting
    FRIENDLY_NAME=image_compare

    if [[ ${OS} == 'windows' ]]; then
      FRIENDLY_NAME=${FRIENDLY_NAME}.exe
    fi

    mv bin/${BINARY_FILE} bin/${FRIENDLY_NAME}
    
    chmod +x bin/${FRIENDLY_NAME}

    tar cf bin/image_compare-${OS}-${ARCH}.tar -C bin ${FRIENDLY_NAME} --remove-files
  done
done