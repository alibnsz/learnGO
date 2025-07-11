package domain

// Product Product, veritabanındaki ürün bilgilerini temsil eden bir yapıdır.
// Domain katmanında tanımlanmıştır çünkü iş mantığının temel veri modelidir.
// Bu yapı, uygulama içerisinde ürün verilerini taşımak için kullanılır.
type Product struct {
	Id       int64   // Ürünün benzersiz kimlik numarası
	Name     string  // Ürünün adı
	Price    float32 // Ürünün fiyatı
	Discount float32 // Ürüne uygulanan indirim oranı
	Store    string  // Ürünün satıldığı mağaza adı
}
