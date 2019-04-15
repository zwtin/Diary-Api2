package controllers

import (
	"Diary-Api2/models"
)

// User is
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Create ...
func (c User) Create(number int, name string) interface{} {
	repo := models.NewUserRepository()
	user := repo.Create(number, name)
	return user
}

// Read ...
func (c User) Read(number int) interface{} {
	repo := models.NewUserRepository()
	user := repo.Read(number)
	return user
}

// Update ...
func (c User) Update(origin int, changed int, name string) interface{} {
	repo := models.NewUserRepository()
	user := repo.Update(origin, changed, name)
	return user
}

// Delete ...
func (c User) Delete(number int) interface{} {
	repo := models.NewUserRepository()
	user := repo.Delete(number)
	return user
}
