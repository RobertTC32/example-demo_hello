#!/bin/bash
#
# -- this run.sh script is replaced by compose.yaml file;
# -- must port 9004 be made available in firewall for app to run on app server ???
# ~/apps/demo_hello/run.sh

#
# run the demo_hello container on app server
docker run --name demo_hello -d --restart unless-stopped -p 9004:80 --env-file ~/apps/demo_hello/.env demo_hello:latest
