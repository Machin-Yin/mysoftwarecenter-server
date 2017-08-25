#!/bin/bash
O=${1-0}
curl -X PUT http://127.0.0.1:8888/user/$O -d "user_name=一铭$O" | jq .
