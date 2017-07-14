#!/bin/bash
# 产品ID, 版本, 图标, 下载地址, 大小 
curl -H "Content-Type: application/json" -X POST -d '[1, 2, 3, 4,54]' http://127.0.0.1:8888/releases | jq .
