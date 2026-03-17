package models

type Kullanici struct {
	ID           uint   `gorm:"primaryKey"`
	KullaniciAdi string `gorm:"unique;not null"`
	Sifre        string `gorm:"not null"`
}
