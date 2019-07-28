APP_VERSION=1.0.0

GIT_HASH=`git rev-parse HEAD`
BUILD_TIME_UTC=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
GO_VERSION=`go version`

go build -ldflags "-X 'main.goBuildVersion=${GO_VERSION}' -X main.buildTime=${BUILD_TIME_UTC} -X main.gitHash=${GIT_HASH} -X main.version=${APP_VERSION}" image_compare.go
