#!/bin/sh

# set -eux
set -eu

function sendRequest() {
    curl \
      -L \
      -A "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36" \
      -e "https://hacpai.com/activity/checkin" \
      -b "$cookie" \
      -c "cookie.txt" \
      -i \
      -o result.html \
      $1
}

cookie=$HACPAI_COOKIE
sendRequest 'https://hacpai.com/activity/checkin'
# cookie='cookie.txt'
sendRequest $(sed -n "s/.*\(https:\/\/hacpai.com\/activity\/daily-checkin?token=\w*\).*/\1/p" result.html)
message=$(sed -n "s/.*今日签到获得[^>]*>\([0-9]*\).*/\1/p" result.html)

curl https://sc.ftqq.com/$SCKEY.send?text=$(date "+%Y-%m-%d")-Hacpai-${message:-失败}
