#!/bin/sh

set -eux

function sendRequest() {
    curl \
      -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36" \
      -H "Cookie: $STUDYGOLANG_COOKIE" \
      -o result.html \
      -L \
      $1
    cat result.html
}

sendRequest 'https://studygolang.com/mission/daily/redeem'
message=$(sed -n "s/.*\(已成功领取每日登录奖励\).*/\1/p" result.html)

curl https://sc.ftqq.com/$SCKEY.send?text=$(date "+%Y-%m-%d")-Studygolang-${message:-失败}
