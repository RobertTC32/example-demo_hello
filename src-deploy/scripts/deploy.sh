#!/bin/bash
#
# chmod +x ~/deploy-demo_hello/*.sh
# ~/deploy-demo_hello/deploy.sh

# copy and organize files in destination directory on devdep server
rm -rf ~/apps/demo_hello
mkdir -p ~/apps/demo_hello
mv ~/deploy-demo_hello/{.,}* ~/apps/demo_hello
#
# load the docker image on app server
cd ~/apps/demo_hello
docker load -i ~/apps/demo_hello/demo_hello.tar
docker image ls | grep demo_hello
#
# -- make port 9001 available for app in firewall
rm ~/apps/demo_hello/deploy.sh
rm -rf ~/deploy-demo_hello
