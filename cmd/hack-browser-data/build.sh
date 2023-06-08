CGO_ENABLED=1 go build 
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64  go build -o hack-browser-data_amd64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build 
