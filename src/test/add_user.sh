#!/bin/bash
curl -X POST http://127.0.0.1:8888/user -d "user_name=$1&avatar_url=$2&mail=$3" | jq .
