#!bin/bash
docker network create -d bridge my-net
docker run -d --name=redis_1  --network my-net --port 6001:6379 redis
docker run -d --name=redis_2  --network my-net --port 6002:6379 redis

