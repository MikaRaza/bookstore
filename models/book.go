package models

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
	Price  float64
}
