package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sky-vibhu/go-userstore/pkg/config"
)

var db *gorm.DB

type User struct{
	gorm.Model
	Id string `gorm:""json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserbyId(Id int64) (*User, *gorm.DB) {
	var getUser User
	db :=db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}