package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/forgoty/go-rest-example/internal/app/model"
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
			name: "shouldBeValid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "emptyEmail",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalidEmailFormat",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalidemailformat"
				return u
			},
			isValid: false,
		},
		{
			name: "emptyPassword",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalidPassword",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "123"
				return u
			},
			isValid: false,
		},
		{
			name: "withEncryptedPassword",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "123"
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
