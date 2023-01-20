package models

type Product struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"varchar(100)" json:"name"`
	Quantity string `gorm:"varchar(100)" json:"quantity"`
}
