## CTI BLOG - Cyber Threat Intelligence & Analysis Portal

CTI BLOG, siber tehdit istihbaratı verilerini toplamak, analiz etmek ve profesyonel bir arayüzle sunmak amacıyla geliştirilmiş, Go (Gin Gonic) tabanlı bir içerik yönetim sistemidir. Siber güvenlik analistlerinin bulgularını (OSINT, Malware Analysis, Dark Web Research vb.) dokümante etmeleri için optimize edilmiştir.
## Öne Çıkan Özellikler
- Güvenlik Odaklı Mimari

    Secure File Upload: Yüklenen görseller için UUID tabanlı isimlendirme, MIME-type doğrulaması ve 5MB dosya boyutu sınırı ile RCE (Remote Code Execution) koruması.

    Session Management: Middleware tabanlı yetkilendirme ve güvenli session yönetimi.

    Input Sanitization: Veritabanı etkileşimlerinde GORM (ORM) kullanılarak SQL Injection risklerine karşı önlem.

-Teknik Özellikler

    Full CRUD Ops: Blog yazıları için oluşturma, okuma, güncelleme ve silme desteği.

    Analist Notları: Her analiz altına dinamik olarak teknik notlar ekleme ve takip etme sistemi.

    Cyber UI: Siber güvenlik estetiğine uygun, neon mor vurgulu Dark Mode arayüz.

    Responsive Design: Mobil ve masaüstü cihazlar için tam uyumlu grid yapısı.

##  Proje Yapısı

```text
.
├── database/     # Veritabanı bağlantı ve init işlemleri
├── handlers/     # Route handler fonksiyonları (Admin, Blog, Auth)
├── middleware/   # AuthRequired kontrolü ve Session yönetimi
├── models/       # GORM veri modelleri (Post, Comment, User)
├── static/       # CSS, JS ve Yüklenen Görseller (uploads)
├── templates/    # HTML şablonları (Go Templates)
├── main.go       # Uygulama giriş noktası
└── go.mod        # Bağımlılık yönetimi
```
## Kurulum ve Çalıştırma
Gereksinimler

    Go 1.20+

    SQLite

## Adımlar

    Projeyi klonlayın:
    Bash

    git clone https://github.com/ozgeclkk/cti_blog.git
    cd cti-blog

    Bağımlılıkları yükleyin:
    Bash

    go mod tidy

    Uygulamayı başlatın:
    Bash

    go run main.go

    Tarayıcınızda açın: http://localhost:5050
