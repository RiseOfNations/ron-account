GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kada-account main/main.go
scp ./kada-account root@39.106.207.62:/root