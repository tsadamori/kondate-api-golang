#!/bin/bash
USER=$1
PASSWORD=$2
DB_HOST=$3
DB_PORT=$4
DB_NAME=$5

PATH="mysql://${USER}:${PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?ssl=disable"
migrate -path database/migrations -path $PATH

PATH="mysql://${USER}:${PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?ssl=disable"
migrate -path database/migrations -path 'mysql://kondate_user:nA4hgwf3@127.0.0.1:3306/kondate?ssl=disable'