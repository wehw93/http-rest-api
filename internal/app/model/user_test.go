package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u:=model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u:=model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u:=model.TestUser(t)
				u.Password = "pw"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u:=model.TestUser(t)
				u.Password = "wehw93@mail.ru"
				return u
			},
			isValid: true,
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
