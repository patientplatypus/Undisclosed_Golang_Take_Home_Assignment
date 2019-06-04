#!/bin/bash

# remove any previous docker images with name htopimg - check that this isn't shadowed first by running `docker images`
docker rmi htopimg
# building the image
# https://stackoverflow.com/questions/19537645/get-environment-variable-value-in-dockerfile (how to pass arguments to Dockerfile)
docker build -t="htopimg" .
# To run without version: 
docker run -it --rm htopimg
# To run with version
# docker run -it --rm htopimg version


#If you were interested in build time as opposed to run time of the first part of the docker question you can approach it like this: 

# bash script

# # remove any previous docker images with name htopimg - check that this isn't shadowed first by running `docker images`
# docker rmi htopimg
# # building the image
# # https://stackoverflow.com/questions/19537645/get-environment-variable-value-in-dockerfile (how to pass arguments to Dockerfile)
# docker build --build-arg version=run -t="htopimg" .
# # Also make sure that 
# docker run -it --rm htopimg

# dockerfile

# FROM debian:latest
# RUN apt-get update
# RUN apt-get install htop
# # # Since htop is a terminal application we also have to set xterm. See links for more info:
# # # https://stackoverflow.com/questions/27826241/running-nano-in-docker-container
# # # https://github.com/moby/moby/issues/9299
# # # https://en.wikipedia.org/wiki/Xterm
# ENV TERM xterm 
# ARG version=noflagfound
# ENV version=$version
# CMD sh -c 'if [ "$version" = run ]; then htop; elif [ "$version" = show ]; then htop -v; else echo no suitable flag found; fi'