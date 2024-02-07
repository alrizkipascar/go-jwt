# NOTES


## LIBRARY
 github.com/gorilla/mux
 github.com/lib/pq

## DOCKER START
 docker run --name some-postgres -e POSTGRES_PASSWORD=test -p 5432:5432 -d postgres

 <!-- FOR SOME REASON YOU NEED TO CHANGE THE PORT to  connect with pgAdmin -->
docker run --name jwt-postgres -e POSTGRES_PASSWORD=test -p 8001:5432 -d postgres
## POSTGRES CONNECT FROM DOCKER to Local PGAdmin
1. docker ps
2. docker inspect <dockerContainerId> | grep IPAddress
3. get IP
4. Put IP and info into pgAdmin

