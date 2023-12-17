package models

import (
	"dbo/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"size:255;no null;" json:"id"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type AuthenticationInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
