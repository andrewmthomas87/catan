#!/bin/sh

docker exec catan_postgres_1 psql -U postgres -c "CREATE DATABASE catan;"
