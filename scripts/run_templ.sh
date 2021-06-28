#! /bin/bash

# exporting the env vars for the templ microservice
docker-ip() {
  if grep -q microsoft /proc/version; then
    echo "localhost"
  else
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' templ-postgres
  fi
}

export PORT=8081
export DB_PORT=5432
export DB_HOST=$(docker-ip)
export DB_USER=templ_user
export DB_PASSWORD=templ
export DB_NAME=templ
export LOG_LEVEL=debug

# build executable
go build -o templ

# launch the templ microservice executable
./templ
