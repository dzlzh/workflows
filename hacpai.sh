#!/bin/sh

set -eux

function sendRequest() {
    curl \
      -A "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36" \
      -b "$cookie" \
      -c "cookie.txt" \
      -i \
      -o result.html \
      $1
}

cookie=$HACPAI_COOKIE
base_url='https://hacpai.com/activity/checkin'
sendRequest $base_url
cookie='cookie.txt'
sign_url=$(sed -n "s/.*\(https:\/\/hacpai.com\/activity\/daily-checkin?token=\w*\).*/\1/p" result.html)
sendRequest $sign_url
sendRequest $base_url
message=$(sed -n "s/.*今日签到获得[^>]*>\([0-9]*\).*/\1/p" result.html)

curl https://sc.ftqq.com/$SCKEY.send?text=$(date "+%Y-%m-%d")-Hacpai-${message:-失败}
