#!/bin/bash

# chmod +x demo_hello-*.sh
#
# copy tar and script files to destination directory
rm -rf ~/apps/demo_hello
mkdir -p ~/apps/demo_hello
mv ~/demo_hello*.* ~/apps/demo_hello
#
# load the docker image and check
docker load -i .~/apps/demo_hello/demo_hello.tar
docker image ls | grep demo_hello
