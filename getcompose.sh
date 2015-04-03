#!/bin/sh

VERSION=1.1.0

curl -L https://github.com/docker/compose/releases/download/${VERSION}/docker-compose-`uname -s`-`uname -m` > docker-compose
chmod +x docker-compose
