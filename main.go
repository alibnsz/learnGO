package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"learnGo/common/postgresql"
	"learnGo/controller"
	"learnGo/persistence"
	"learnGo/service"
)

// main fonksiyonu, uygulamanın giriş noktasıdır.
// Bu fonksiyon çalıştırıldığında web sunucusu başlatılır.
func main() {
	// Echo web framework'ünün yeni bir örneğini oluşturuyoruz
	// Echo, Go dilinde hızlı ve minimal bir web framework'tür
	e := echo.New()
	
	// Veritabanı bağlantısı oluştur
	dbPool := postgresql.GetConnectionPool(context.Background(), postgresql.Config{
		Host:                   "localhost", 
		Port:                   "5432",      // PostgreSQL varsayılan portu
		DbName:                 "learnGo",
		UserName:               "postgres",
		Password:               "postgres",
		MaxConnections:         "10",
		MaxConnectionsIdleTime: "30s",
	})
	
	// Dependency Injection ile uygulama bileşenlerini oluştur
	// Repository, Service ve Controller katmanları birbirine bağlanır
	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	
	// Kontrolcü rotalarını Echo'ya kaydet
	productController.RegisterRoutes(e)
	
	// Web sunucusunu localhost:8080 adresinde başlatıyoruz
	// Bu adres üzerinden HTTP istekleri kabul edilecektir
	log.Info("Server starting at localhost:8080")
	if err := e.Start("localhost:8080"); err != nil {
		log.Error("Server failed to start: ", err)
	}
}
