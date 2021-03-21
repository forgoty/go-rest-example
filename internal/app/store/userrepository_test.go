package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/forgoty/go-rest-example/internal/app/model"
	"github.com/forgoty/go-rest-example/internal/app/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmailWhenLookupFailed(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().FindByEmail(model.TestEmail())
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().Create(model.TestUser(t))
	if err != nil {
		t.Fatal(err)
	}

	u, err := s.User().FindByEmail(model.TestEmail())

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
