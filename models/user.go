package models

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email" gorm:"unique"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	ProfileImage   string    `json:"profile_image"`
	HashedPassword string    `json:"-"`
	JoinedAt       time.Time `json:"joined_at" gorm:"autoCreateTime"`
	LastLogin      time.Time `json:"last_login" gorm:"null"`
}

func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ValidatePhone(phone string) bool {
	reg, _ := regexp.Compile("0?9[0-9]{9}")
	return reg.MatchString(phone)
}
