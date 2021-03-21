package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/forgoty/go-rest-example/internal/app/model"
	"github.com/forgoty/go-rest-example/internal/app/store"
)

var (
	testEmail = "test@test.com"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: testEmail,
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmailWhenLookupFailed(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().FindByEmail(testEmail)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().Create(&model.User{
		Email: testEmail,
	})
	if err != nil {
		t.Fatal(err)
	}

	u, err := s.User().FindByEmail(testEmail)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
