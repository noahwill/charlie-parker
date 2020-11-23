#!/bin/bash

run_type=$1

case $run_type in 
  local)
    echo "Turning on application in local mode"
    docker rm cp-dynamo
    docker rm cp-seeder
    docker rm cp-server-local
    docker-compose -f docker-compose.local.yml down
    docker-compose -f docker-compose.local.yml rm -f
    docker-compose -f docker-compose.local.yml build
    docker-compose -f docker-compose.local.yml up
    ;;
esac