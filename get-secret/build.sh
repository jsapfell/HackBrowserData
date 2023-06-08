CGO_ENABLED=1 go build  -o Keychain
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64  go build  -o Keychain_amd64
