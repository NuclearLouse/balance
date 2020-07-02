package models_test

import (
	"balance/internal/app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypt password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				u.HashPassword = "hashpassword"
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = "pas"
				return u
			},
			isValid: false,
		},
		{
			name: "short repeat password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.RepeatPassword = "pas"
				return u
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_PasswordHashing(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.PasswordHashing())
	assert.NotEmpty(t, u.HashPassword)
}