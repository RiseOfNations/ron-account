GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ron-account ../main/main.go
scp ./ron-account root@39.106.207.62:/root
rm ron-account