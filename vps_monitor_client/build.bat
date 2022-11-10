SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOPROXY=https://goproxy.cn
go build -o vps-client main.go