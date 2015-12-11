#!/bin/bash
docker build -t zaddone/dimon .
docker run --link redis:redis --link postgresql:postgresql --link mysql:mysql -it --rm -p 8080:8080 --name dimon_site  zaddone/dimon
