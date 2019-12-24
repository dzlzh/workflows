#!/bin/sh

set -eux

base_url='https://www.v2ex.com/mission/daily/'

function sendRequest() {
    curl \
      -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36" \
      -H "Cookie: $V2EX_COOKIE" \
      -o result.html \
      $1
}

sendrequest $base_url
sign_url=$base_url`sed -n "s/.*\/mission\/daily\/\(redeem?once=\w*\).*/\1/p" result.html`
sendrequest $sign_url
sendrequest $base_url
message=$(sed -n "s/.*\(每日登录奖励已领取\).*/\1/p" result.html)
curl https://sc.ftqq.com/$SCKEY.send?text=$(date "+%Y-%m-%d")-V2ex-${message:-失败}
