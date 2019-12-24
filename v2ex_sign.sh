#!/bin/sh

set -eux

UA="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"
BASE_URL='https://www.v2ex.com/mission/daily/'

function sendRequest() {
    curl \
      -H "User-Agent: $UA" \
      -H "Cookie: $V2EX_COOKIE" \
      -o result.html \
      $1
}

sendRequest $BASE_URL

SIGN_URL=$BASE_URL`sed -n "s/.*\/mission\/daily\/\(redeem?once=\w*\).*/\1/p" result.html`
sendRequest $SIGN_URL

sendRequest $BASE_URL
message=$(sed -n "s/.*\(每日登录奖励已领取\).*/\1/p" result.html)
curl \
    -s --user "api:$MAIL_API_KEY" \
    https://api.mailgun.net/v3/$MAIL_DOMAIN/messages \
    -F from="Notification <mailgun@$MAIL_DOMAIN>" \
    -F to=$MAIL_TO \
    -F subject="$(date "+v2ex %Y-%m-%d")" \
    -F text="${message:-失败}"
