package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Password       string    `json:"password,omitempty"`
	RepeatPassword string    `json:"repeat_password,omitempty"`
	HashPassword   string    `json:"-"`
	Username       string    `json:"username,omitempty"`
	Admin          bool      `json:"admin"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	Comment        string    `json:"comment,omitempty"`
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
	u.RepeatPassword = ""
}

// Validate ... проверка введенных данных при логине
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.HashPassword == "")), validation.Length(4, 100)),
		validation.Field(&u.RepeatPassword, validation.By(requiredIf(u.HashPassword == "")), validation.Length(4, 100)),
	)
}

// PasswordHashing ...
func (u *User) PasswordHashing() error {
	if len(u.Password) != 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.HashPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ComparePassword ...
func (u *User) ComparePassword(password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(password)) == nil
}
