#!/bin/bash
O=${1-16}
curl http://127.0.0.1:8888/banner/$O | jq .
