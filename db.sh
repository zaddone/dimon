#!/bin/bash
docker stop postgresql redis mysql
docker rm postgresql redis mysql
docker run -d -p 5432:5432 --name postgresql -e POSTGRES_PASSWORD=test2015 -e PGDATA=/data -e POSTGRES_DB=dimon  -v ~/data:/data postgres
docker run --name redis -p 6379:6379 -d redis
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=test2015 -e MYSQL_DATABASE=dimon -v ~/data_mysql:/var/lib/mysql -p 3306:3306 mysql:latest
