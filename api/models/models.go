package models

type ShortUrl struct {
	Url      string `json:"url" gorm:"not null"`
	ShortUrl string `json:"shorturl" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}
