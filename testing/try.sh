#!/usr/bin/env bash

readonly PROLAMB_URL=""

TARGET=`base64 -w0 - << EOF
{
  "name": "test",
  "type": "http",
  "options": {
    "method": "GET",
    "url": "https://google.com"
  }
}
EOF`

echo "${PROLAMB_URL}?target=$TARGET"
