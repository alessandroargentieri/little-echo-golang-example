#! /bin/bash

if [[ $( docker images | grep postgres) ]]; then
    echo "postgres image already pulled."
else
	echo "pulling postgres docker image"
	docker pull postgres    
fi

if [[ $( docker ps -a | grep templ-postgres | head -c12) ]]; then
    echo "postgres container already running."
else
	docker run --name templ-postgres -e POSTGRES_PASSWORD=templ -p 5432:5432 -d postgres 
fi

export CONTAINER_ID=$( docker ps -a | grep templ-postgres | head -c12)

sleep 2s
docker exec -it $CONTAINER_ID psql -U postgres -c "create user templ_user with password 'templ';"
sleep 0.2
docker exec -it $CONTAINER_ID psql -U postgres -c "create database templ owner templ_user;"
sleep 0.2
docker exec -it $CONTAINER_ID psql -U postgres -c "grant all privileges on database templ to templ_user;"

sleep 0.2
docker exec -it $CONTAINER_ID psql -U templ_user -d templ -c "CREATE TABLE IF NOT EXISTS users (id varchar(50) CONSTRAINT user_pk PRIMARY KEY, name varchar(50) NOT NULL, surname varchar(50) NOT NULL, created_at timestamp NOT NULL, updated_at timestamp NOT NULL);"

# restart container
docker start $CONTAINER_ID




