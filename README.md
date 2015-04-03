# DOcker Survival Kit

This repo contains utilities that make stuff easier with Docker, especially on CoreOS.

To use this, just put this directory to your PATH.

## Build

```
$ go build dops.go -o dops
$ sh getcompose.se
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

- It's too wide - the listing lines span multiple lines in my 80-char terminal window
- It's difficult to read - it takes time to find the ID of a container of the image you want. 
- It doesn't show IP addresses of the containers.

I wrote this simple tool in Go which prints running containers in colors.

## denter

If you have running container and you want to have a shell in it for checking stuff, then you might appreciate this script which allows to connect to a container by mathcing substring in docker ps output.

If you run it witouht parameters, you will get shell in the latest run container.

![How it looks](http://i.imgur.com/wyjHrrP.png)

