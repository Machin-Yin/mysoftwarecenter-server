#!/bin/bash
# 评论信息ID, 产品ID(sc_product), 发布ID(sc_release), 用户ID, 评论内容，具体评分，评论时间 
product_ID=`curl -G "http://127.0.0.1:8888/product_from_name" --data-urlencode "product_name=$1" | jq '.product[].ID'`
release_ID=`curl -G "http://127.0.0.1:8888/product_from_name" --data-urlencode "product_name=$1" | jq '.product[].release_ID'`
[[ -z "$release_ID" ]] && echo "------------------->" $1
[[ -z "$product_ID" ]] && echo "------------------->" $1
curl -X POST http://127.0.0.1:8888/comment -d "product_ID=$product_ID&release_ID=$release_ID&user_ID=$4&comment_text=$2&comment_grade=$3&comment_date=$4" | jq .
#curl -X POST http://127.0.0.1:8888/release -d "ID=$2&product_ID=$product_ID&release_ID=$release_ID&user_ID=$3&comment_text=$4&comment_grade=$5&comment_date=$6" | jq .
#echo curl -X POST http://127.0.0.1:8888/release -d "product_ID=$product_ID&version=$2&icon_url=$3&download_url=$4&changelog=Change Log&package_size=$5&package_type=1&release_grade=1" 
