## CTI BLOG - Cyber Threat Intelligence & Analysis Portal
<img width="1837" height="883" alt="resim" src="https://github.com/user-attachments/assets/0455b0e7-b5aa-409c-bfb9-0fc501612d2e" />


CTI BLOG, siber tehdit istihbaratı verilerini toplamak, analiz etmek ve profesyonel bir arayüzle sunmak amacıyla geliştirilmiş, Go (Gin Gonic) tabanlı bir içerik yönetim sistemidir. Siber güvenlik analistlerinin bulgularını (OSINT, Malware Analysis, Dark Web Research vb.) dokümante etmeleri için optimize edilmiştir.
## Öne Çıkan Özellikler
<img width="967" height="607" alt="resim" src="https://github.com/user-attachments/assets/e1b61b7d-7eef-4810-bd5e-eb49e92146e0" />
<img width="1837" height="889" alt="resim" src="https://github.com/user-attachments/assets/c14c2868-5ce0-476a-9d02-2008bd89be65" />


- Güvenlik Odaklı Mimari



    Secure File Upload: Yüklenen görseller için UUID tabanlı isimlendirme, MIME-type doğrulaması ve 5MB dosya boyutu sınırı ile RCE (Remote Code Execution) koruması.

    Session Management: Middleware tabanlı yetkilendirme ve güvenli session yönetimi.

    Input Sanitization: Veritabanı etkileşimlerinde GORM (ORM) kullanılarak SQL Injection risklerine karşı önlem.

-Teknik Özellikler
<img width="1284" height="681" alt="resim" src="https://github.com/user-attachments/assets/03809849-a722-474e-94f4-1516e183ca74" />

<img width="952" height="727" alt="resim" src="https://github.com/user-attachments/assets/98094fa7-c7f0-40af-89d1-e290850ab3e0" />



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
