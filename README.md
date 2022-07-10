## usage:
```
./sshd-linux-aarch64 
2022/07/10 05:12:18 starting ssh server on port 2222... tty.go
```

## connect:
`ssh -o HostKeyAlgorithms=+ssh-rsa -o PubkeyAcceptedKeyTypes=+ssh-rsa 127.0.0.1 -p 2222`
#### default password: 	`qgb`	

## build: 
```
$ ./build.sh 
++ go get
++ export CGO_ENABLED=0 GOOS=linux GOARCH=arm64
++ CGO_ENABLED=0
++ GOOS=linux
++ GOARCH=arm64
++ go build -trimpath -ldflags '-s -w -buildid=' -o sshd-linux-arm64
++ export CGO_ENABLED=0 GOOS=linux GOARCH=amd64
++ CGO_ENABLED=0
++ GOOS=linux
++ GOARCH=amd64
++ go build -trimpath -ldflags '-s -w -buildid=' -o sshd-linux-amd64


```

## download:
 https://github.com/qgb/sshd/releases
