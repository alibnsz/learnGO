package controller

import (
	"github.com/labstack/echo/v4"
	"learnGo/service"
	"net/http"
)

// ProductController, ürünlerle ilgili HTTP isteklerini karşılayan ve işleyen kontrolcüdür.
// Bu kontrolcü, gelen HTTP isteklerini uygun servis çağrılarına dönüştürür.
type ProductController struct {
	productService service.IProductService // Ürün servisi
}

// NewProductController, yeni bir ProductController örneği oluşturur.
// Dependency Injection prensibi ile dışarıdan bir servis alır.
func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// RegisterRoutes, ürün ile ilgili rotaları Echo router'ına kaydeder.
// Bu metot, kontrolcünün hangi HTTP yollarını (endpoint) dinleyeceğini belirler.
func (pc *ProductController) RegisterRoutes(e *echo.Echo) {
	// Ürünlerle ilgili rotaları "/products" ön eki altında gruplandır
	productGroup := e.Group("/products")
	
	// GET /products - Tüm ürünleri getir
	productGroup.GET("", pc.GetAllProducts)
	
	// İleride eklenecek diğer rotalar:
	// productGroup.GET("/:id", pc.GetProductById)
	// productGroup.POST("", pc.CreateProduct)
	// productGroup.PUT("/:id", pc.UpdateProduct)
	// productGroup.DELETE("/:id", pc.DeleteProduct)
}

// GetAllProducts, tüm ürünleri getiren HTTP handler fonksiyonudur.
// Bu metot, HTTP GET /products isteğini karşılar ve tüm ürünleri JSON formatında döndürür.
func (pc *ProductController) GetAllProducts(c echo.Context) error {
	// Servis katmanından tüm ürünleri al
	products := pc.productService.GetAllProducts()
	
	// Ürünleri JSON formatında HTTP yanıtı olarak döndür
	return c.JSON(http.StatusOK, products)
} 