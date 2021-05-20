GOOS=linux go build -o bin/WpGo_linux main.go
GOOS=windows go build -o bin/WpGo.exe main.go
upx -9 bin/*
rm -rf main
