package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

var (
	db * gorm .DB
)

func Connect() {
	d, err := gorm.Open("mysql", "shivam:shivamyadav@12@/simplerest")
	if err != nil{
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB{
	return db
}