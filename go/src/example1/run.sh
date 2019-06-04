#!/bin/bash

docker rmi $(docker images -q)
docker build -t example1 .
docker run -it -p 8000:8000 example1