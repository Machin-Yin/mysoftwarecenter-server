#!/bin/bash
# 产品ID, 优先级
product_ID=`curl -G "http://127.0.0.1:8888/product_from_name" --data-urlencode "product_name=$1" | jq '.product[].ID'`
[[ -z "$product_ID" ]] && echo "------------------->" $1
curl -X POST http://127.0.0.1:8888/banner -d "ID=$product_ID&priority=$2" | jq .
