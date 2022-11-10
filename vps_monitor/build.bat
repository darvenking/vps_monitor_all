SET GOPROXY=https://goproxy.cn,direct
gf build main.go -n vps-server -a arm -s linux -p .
