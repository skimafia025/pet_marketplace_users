package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"UUID" gorm:"primaryKey"`
	EMAIL     string    `json:"EMAIL" gorm:"unique;not null"`
	PASSWORD  string    `json:"PASSWORD" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Email    string `json:"EMAIl" binding:"required,email"`
	Password string `json:"PASSWORD" binding:"required,min=8"`
	Name     string `json:"NAME" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required"`
}

func (u *User) Hashpassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PASSWORD), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PASSWORD = string(hashedPassword)
	return nil
}
func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PASSWORD), []byte(password)) == nil
}
