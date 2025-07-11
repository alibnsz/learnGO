package postgresql

// Config Config, PostgreSQL veritabanı bağlantısı için gerekli yapılandırma parametrelerini içeren yapıdır.
// Bu yapı sayesinde veritabanı bağlantı bilgileri tek bir yerden yönetilebilir.
// Farklı ortamlarda (geliştirme, test, üretim) farklı konfigürasyonlar kullanılabilir.
type Config struct {
	Host                   string // Veritabanı sunucusunun adresi (örn. localhost)
	Port                   string // Veritabanı sunucusunun portu (örn. 5432)
	UserName               string // Veritabanı kullanıcı adı
	Password               string // Veritabanı şifresi
	DbName                 string // Bağlanılacak veritabanının adı
	MaxConnections         string // Maksimum eşzamanlı bağlantı sayısı
	MaxConnectionsIdleTime string // Kullanılmayan bağlantıların kapatılmadan önce bekleyeceği süre
}
