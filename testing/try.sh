#!/bin/bash

readonly PROLAMB_URL=""

TARGET=`cat << EOF | base64
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
