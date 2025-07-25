#!/bin/bash

# Bu script, testler için gerekli olan PostgreSQL veritabanını Docker üzerinde oluşturur.
# Testler sırasında gerçek veritabanı yerine bu test veritabanı kullanılır.

# 1. PostgreSQL container başlatılıyor
# Docker kullanarak izole bir PostgreSQL sunucusu oluşturulur
# Bu sayede testler gerçek veritabanını etkilemez
docker run --name postgres-test \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 6432:5432 \
  -d postgres:latest

echo "PostgreSQL starting..."
sleep 3  # Container'ın başlaması için bekle

# 2. learnGo veritabanı oluşturuluyor
# Uygulamanın kullanacağı veritabanı oluşturulur
docker exec -i postgres-test psql -U postgres -d postgres -c "CREATE DATABASE \"learnGo\";"
echo "Database learnGo created."
sleep 1  # İşlemin tamamlanması için kısa bir bekleme

# 3. products tablosu oluşturuluyor
# Uygulama için gerekli olan tablo yapısı oluşturulur
docker exec -i postgres-test psql -U postgres -d learnGo -c "
CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DOUBLE PRECISION NOT NULL,
  discount DOUBLE PRECISION,
  store VARCHAR(255) NOT NULL
);"
echo "Table products created."
