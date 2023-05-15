set -x  # ECHO on
#go get github.com/google/shlex github.com/gliderlabs/ssh
export GOPROXY=https://goproxy.cn
go get

export BUILD_COMMAND='go build -trimpath -ldflags "-s -w -buildid=" -buildvcs=false -o'
#golang 1.20  -buildvcs=false 

export CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ;eval $BUILD_COMMAND sshd-$GOOS-$GOARCH
export CGO_ENABLED=0 GOOS=linux GOARCH=arm   ;eval $BUILD_COMMAND sshd-$GOOS-$GOARCH
export CGO_ENABLED=0 GOOS=linux GOARCH=arm64 ;eval $BUILD_COMMAND sshd-$GOOS-$GOARCH