#!/bin/bash

# 私有化环境变量，服务中依赖该变量做私有化变量的控制
export scenes="privatization"

cd /usr/local/app/bin
./helloworld 2>&1 | tee /usr/local/app/log/server.log
