# !/bin/bash
# How to run: sh publish.sh 1.0

docker build -t prongbang/wiremock:$0 .
docker tag wiremock:$0 prongbang/wiremock:$0
docker push prongbang/wiremock:$0