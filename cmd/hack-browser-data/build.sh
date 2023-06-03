CGO_ENABLED=1 go build -o Chrome
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o Chrome.exe
