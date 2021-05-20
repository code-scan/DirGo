GOOS=linux go build -o bin/DirGO main.go
GOOS=windows go build -o bin/DirGO.exe main.go
upx -9 bin/*
rm -rf main
