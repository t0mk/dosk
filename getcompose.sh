#!/bin/sh

VERSION=1.1.0
OS=darwin
ARCH=x86_64

curl -L https://github.com/docker/compose/releases/download/${VERSION}/docker-compose-${OS}-${ARCH} > docker-compose

chmod +x docker-compose
