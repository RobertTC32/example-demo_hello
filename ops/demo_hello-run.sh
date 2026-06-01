#!/bin/bash

# make port 9001 available in firewall
#
# run the demo_hello container
docker run  --name demo_hello --rm -p 9001:8080 -e OCI_PORT=9001 -e LOG_LEVEL=info demo_hello:latest
