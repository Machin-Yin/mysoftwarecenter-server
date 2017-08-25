#!/bin/bash
# 评论信息ID, 产品ID(sc_product), 发布ID(sc_release), 用户ID, 评论内容，具体评分，评论时间 
curl -H "Content-Type: application/json" -X POST -d '[1, 2, 3, 4,54]' http://127.0.0.1:8888/comments | jq .
