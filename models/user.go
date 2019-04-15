package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "hoge"
	PROTOCOL := ""
	DBNAME := "bookshelf"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// User is
type User struct {
	gorm.Model
	Number int
	Name   string
}

// NewUser ...
func NewUser(number int, name string) User {
	return User{
		Number: number,
		Name:   name,
	}
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// Create ...
func (m UserRepository) Create(number int, name string) *User {
	db := gormConnect()
	defer db.Close()

	user := User{Number: number, Name: name}
	db.Create(&user)
	return &user
}

// Read ...
func (m UserRepository) Read(number int) *User {
	db := gormConnect()
	defer db.Close()

	user := User{}
	db.Where("number = ?", number).First(&user)
	return &user
}

// Update ...
func (m UserRepository) Update(origin int, changed int, name string) *User {
	db := gormConnect()
	defer db.Close()

	user := User{}
	db.Where("number = ?", origin).First(&user)
	user.Number = changed
	user.Name = name
	db.Save(&user)
	return &user
}

// Delete ...
func (m UserRepository) Delete(number int) *User {
	db := gormConnect()
	defer db.Close()

	user := User{}
	db.Where("number = ?", number).First(&user)
	db.Delete(&user)
	return &user
}
