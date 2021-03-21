package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    TestEmail(),
		Password: "password",
	}
}

func TestEmail() string {
	return "test@test.com"
}
