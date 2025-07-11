package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// GetConnectionPool GetConnectionPool, PostgreSQL veritabanına bağlantı havuzu oluşturan fonksiyondur.
// pgxpool kütüphanesi kullanılarak veritabanı bağlantı havuzu oluşturulur.
// Bağlantı havuzu, veritabanı işlemleri için birden fazla bağlantıyı yönetir ve performansı artırır.
//
// Parametreler:
// - context: İşlem bağlamı, bağlantının iptal edilmesi veya zaman aşımı için kullanılır
// - config: Veritabanı bağlantı parametrelerini içeren yapılandırma nesnesi
//
// Dönüş:
// - *pgxpool.Pool: Veritabanı bağlantı havuzu
func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	// Bağlantı dizesi oluşturulur
	// Bu dize, PostgreSQL sunucusuna bağlanmak için gerekli tüm bilgileri içerir
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionsIdleTime)

	// Bağlantı dizesinden konfigürasyon nesnesi oluşturulur
	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		// Konfigürasyon hatası durumunda uygulama durdurulur
		panic(parseConfigErr)
	}

	// Veritabanına bağlantı sağlanır
	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		// Bağlantı hatası durumunda hata loglanır ve uygulama durdurulur
		log.Error("Unable to connect to database: %v\n", err)
		panic(err)
	}

	// Oluşturulan bağlantı havuzu döndürülür
	return conn
}
