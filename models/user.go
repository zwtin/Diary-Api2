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

// GetByID ...
func (m UserRepository) GetByID(number int) *User {
	db := gormConnect()
	defer db.Close()

	user := User{}
	user.Number = number
	db.First(&user)
	return &user
}
