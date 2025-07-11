#!/bin/bash

# 1. PostgreSQL container başlatılıyor
docker run --name postgres-test \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 6432:5432 \
  -d postgres:latest

echo "PostgreSQL starting..."
sleep 3

# 2. learnGo veritabanı oluşturuluyor
docker exec -i postgres-test psql -U postgres -d postgres -c "CREATE DATABASE \"learnGo\";"
echo "Database learnGo created."
sleep 1

# 3. products tablosu oluşturuluyor
docker exec -i postgres-test psql -U postgres -d learnGo -c "
CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DOUBLE PRECISION NOT NULL,
  discount DOUBLE PRECISION,
  store VARCHAR(255) NOT NULL
);"
echo "Table products created."
