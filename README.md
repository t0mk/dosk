# DOcker Survival Kit

This repo contains utilities that make stuff easier with Docker, especially on CoreOS.

To use this, just put this directory to your PATH.

## Build

```
$ GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o dops dops.go
$ ./getcompose.sh
```

## Install

Just put it to PATH, so sth like this from you home dir:

```
git clone https://github.com/t0mk/dosk
[ -L .bashrc ] && rm .bashrc
echo 'export PATH=${PATH}:${HOME}/dosk' >> .bashrc
```

## dops

dops is more pretty listing of runnign containers.

docker ps is ugly:

* It's too wide - the listing lines span multiple lines in my 80-char terminal window
* It's difficult to read - it takes time to find the ID of a container of the image you want. 
* It doesn't show IP addresses of the containers.

I wrote this simple tool in Go which prints running containers in colors.

![How it looks](http://i.imgur.com/wyjHrrP.png)

## denter

If you have running container and you want to have a shell in it for checking stuff, then you might appreciate this script which allows to connect to a container by mathcing substring in docker ps output.

E.g. `denter fe` will open you an interactive bash in a container with ID fe[...]. If you run only `denter`, you will get bash in the latest run container. If you have several containers running and only one of them has name or image called "drupal", you can connect to it by `denter drup`.

## dlog

This is `tail -f` on a container specified by fuzzy name like in denet.
