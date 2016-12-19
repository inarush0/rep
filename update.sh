#!/bin/bash

docker pull $1/rep-prod:latest
if docker stop rep-app; then docker rm rep-app; fi
docker run -d -p 80:8080 --name rep-app $1/rep-prod
if docker rmi $(docker images --filter "dangling=true" -q --no-trunc); then :; fi
