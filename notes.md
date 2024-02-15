# NOTES

There is something wrong with docker and main.go, I cannot move it to cmd

## LIBRARY
 github.com/gorilla/mux
 github.com/lib/pq

## DOCKER START
### REMINDER FOR POSTGRE =  DO NOT USE POSTGRES IN DOCKER FOR PRODUCTION ONLY IN STAGING
 docker run --name some-postgres -e POSTGRES_PASSWORD=test -p 5432:5432 -d postgres

 <!-- FOR SOME REASON YOU NEED TO CHANGE THE PORT to  connect with pgAdmin -->
docker run --name jwt-postgres -e POSTGRES_PASSWORD=test -p 8001:5432 -d postgres

<!-- Start Docker -->
docker start <CONTAINER_ID>

## POSTGRES CONNECT FROM DOCKER to Local PGAdmin
1. docker ps
2. docker inspect <dockerContainerId> | grep IPAddress
3. get IP
4. Put IP and info into pgAdmin

## Other Command

export JWT_SECRET=hunter9999

## SEEDING COMMAND
------> ./bin/go-jwt --seed

