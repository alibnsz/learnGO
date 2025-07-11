package persistence

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"learnGo/domain"
)

// IProductRepository, ürünlerle ilgili veritabanı işlemlerini tanımlayan bir arayüzdür.
// Bu arayüz sayesinde farklı implementasyonlar kullanılabilir ve test edilebilirlik artar.
type IProductRepository interface {
	GetAllProducts() []domain.Product
}

// ProductRepository, IProductRepository arayüzünü uygulayan ve PostgreSQL 
// veritabanı ile iletişim kuran yapıdır.
type ProductRepository struct {
	dbPool *pgxpool.Pool // PostgreSQL bağlantı havuzu
}

// NewProductRepository, yeni bir ProductRepository örneği oluşturur.
// Dependency Injection prensibi ile dışarıdan bir veritabanı bağlantı havuzu alır.
// Bu sayede test edilebilirlik artar ve bağımlılıklar daha iyi yönetilir.
func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

// GetAllProducts, veritabanındaki tüm ürünleri getirir.
// Bu metot, products tablosundaki tüm kayıtları sorgular ve domain.Product listesine dönüştürür.
func (productRepository *ProductRepository) GetAllProducts() []domain.Product {

	ctx := context.Background()
	// Tüm ürünleri seçen SQL sorgusu çalıştırılır
	productRows, err := productRepository.dbPool.Query(ctx, "select * from products")

	if err != nil {
		// Hata durumunda log kaydı oluşturulur ve boş bir dizi döndürülür
		log.Errorf("Error while getting product rows: %v", err)
		return []domain.Product{}
	}

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	// Sorgu sonucunda dönen her satır için Product nesnesi oluşturulur
	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
