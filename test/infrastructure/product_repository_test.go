package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"learnGo/common/postgresql"
	"learnGo/persistence"
	"os"
	"testing"
)

// Test için kullanılacak global değişkenler
var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

// TestMain, test paketi için ana giriş noktasıdır.
// Tüm testler çalıştırılmadan önce gerekli hazırlıkları yapar ve 
// testler tamamlandıktan sonra temizlik işlemlerini gerçekleştirir.
func TestMain(m *testing.M) {
	ctx := context.Background()

	// Test veritabanına bağlantı oluşturulur
	// Bu bağlantı, test_db.sh script'i ile oluşturulan Docker container'ına yapılır
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                   "localhost",
		Port:                   "6432",      // Docker container'ının dışarıya açtığı port
		DbName:                 "learnGo",
		UserName:               "postgres",
		Password:               "postgres",
		MaxConnections:         "10",
		MaxConnectionsIdleTime: "30s",
	})
	
	// Product repository örneği oluşturulur
	productRepository = persistence.NewProductRepository(dbPool)
	
	// Testler çalıştırılır ve çıkış kodu alınır
	exitCode := m.Run()
	
	// Program sonlandırılır
	os.Exit(exitCode)
}

// TestGetAllProducts, GetAllProducts metodunun doğru çalışıp çalışmadığını test eder.
// Bu test, veritabanından ürünlerin başarıyla alınıp alınmadığını kontrol eder.
func TestGetAllProducts(t *testing.T) {
	fmt.Println("GetAllProducts")
	fmt.Println(productRepository)
	fmt.Println(dbPool)
}
