# learnGo

Go dilinde geliştirilmiş basit bir web uygulaması. Bu proje, Echo framework ve PostgreSQL veritabanı kullanılarak oluşturulmuştur.

## Proje Yapısı

Proje, Clean Architecture prensiplerine göre organize edilmiştir:

- `common`: Veritabanı bağlantısı gibi ortak kullanılan bileşenler
- `controller`: HTTP isteklerini karşılayan API kontrolcüleri
- `domain`: Temel veri modelleri ve iş mantığı
- `persistence`: Veritabanı işlemleri ve repository implementasyonları
- `service`: İş mantığı ve servis katmanı
- `test`: Test kodları ve test altyapısı

## Başlangıç

Projeyi çalıştırmak için:

1. PostgreSQL veritabanını kurun
2. Test veritabanını oluşturmak için `test/scripts/test_db.sh` scriptini çalıştırın
3. `go run main.go` komutu ile uygulamayı başlatın

## Teknolojiler

- Go 1.24
- Echo Web Framework
- PostgreSQL
- pgx - PostgreSQL sürücüsü
