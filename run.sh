#!/bin/bash
docker build -t zaddone/dimon .
docker run -v ~/data_sqlite:/data_sqlite -it --rm -p 80:8080 --name dimon_site  zaddone/dimon
