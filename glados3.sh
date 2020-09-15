#!/bin/sh

# set -eux
set -eu

curl 'https://glados.work/api/user/checkin' \
  -H 'authority: glados.work' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'dnt: 1' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36' \
  -H 'content-type: application/json;charset=UTF-8' \
  -H 'origin: https://glados.work' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://glados.work/console/checkin' \
  -H 'accept-language: en,en-US;q=0.9,zh;q=0.8,zh-CN;q=0.7,zh-HK;q=0.6' \
  -H "cookie: ${GLADOS_COOKIE_3}" \
  --data-binary '{}' \
  --compressed \
  -o result.html

message=$(sed -n "s/.*\"message\":\"\([^\"]*\)\".*/\1/p" result.html)
curl -G --data-urlencode "text=$(date '+%Y-%m-%d')-GLaDOS-${message:-失败}" https://sc.ftqq.com/$SCKEY.send
