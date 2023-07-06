package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// swagger:model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
}

/*BeforeCreate*/
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u.Password = string(hashedPassword)

	return
}

/*BeforeSave*/
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()

	if tx.Statement.Changed("Password") {
		hashedPassword, error := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if error != nil {
			err = error
			return
		}

		u.Password = string(hashedPassword)
	}

	return
}

/*CheckPassword*/
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
