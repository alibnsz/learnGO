package service

import (
	"learnGo/domain"
	"learnGo/persistence"
)

// IProductService, ürünlerle ilgili iş mantığı operasyonlarını tanımlayan arayüzdür.
// Bu arayüz, controller katmanı ile persistence katmanı arasında bir köprü görevi görür.
type IProductService interface {
	GetAllProducts() []domain.Product
	// İleride eklenecek diğer iş mantığı metodları burada tanımlanabilir
	// Örneğin: GetProductById, CreateProduct, UpdateProduct, DeleteProduct vb.
}

// ProductService, IProductService arayüzünü uygulayan ve ürünlerle ilgili
// iş mantığı operasyonlarını gerçekleştiren servistir.
type ProductService struct {
	productRepository persistence.IProductRepository // Veritabanı işlemleri için repository
}

// NewProductService, yeni bir ProductService örneği oluşturur.
// Dependency Injection prensibi ile dışarıdan bir repository alır.
func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

// GetAllProducts, tüm ürünleri getiren servis metodudur.
// Bu metot, repository katmanını kullanarak veritabanından ürünleri çeker.
// İleride, ürünlerin filtrelenmesi veya sıralanması gibi iş mantığı eklenebilir.
func (s *ProductService) GetAllProducts() []domain.Product {
	// Şu an için sadece repository'den veri çekiliyor
	// Ancak ileride burada ek iş mantığı uygulanabilir
	return s.productRepository.GetAllProducts()
} 