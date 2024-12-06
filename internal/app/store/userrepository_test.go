package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store"
)

func TestUserReposytoryCreate(t *testing.T) {
	s, teardown := store.TestStore(t, database_url)
	defer teardown("users")
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestFindUserByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, database_url)
	defer teardown("users")
	_, err := s.User().FindByEmail("user@example.org")
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail("user@example.org")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
