#!/bin/bash
docker build -t zaddone/dimon .
docker run -it -p 8080:8080 --name dimon_site --rm zaddone/dimon
