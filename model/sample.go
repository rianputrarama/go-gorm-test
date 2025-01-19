package model

type Sample struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
