#!/bin/bash
curl -H "Content-Type: application/x-www-form-urlencoded"  -X POST http://127.0.0.1:8888/product -d "product_name=$1&product_description=$2&category_ID=$3&vendor_name=$4&url=$5&grade_count=$6&executable_file=$7&icon_url=$8&package_name=$9" | jq .
