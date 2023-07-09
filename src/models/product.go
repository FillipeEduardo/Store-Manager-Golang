package models

type Product struct {
	Id   int32  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(30)"`
}
