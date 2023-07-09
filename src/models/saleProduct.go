package models

type SaleProduct struct {
	ProductId int32   `gorm:"primaryKey"`
	Products  Product `gorm:"foreignKey:ProductId"`
	SaleId    int32   `gorm:"primaryKey"`
	Sales     Sale    `gorm:"foreignKey:SaleId"`
	Quantity  int32
}
