package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store"
	"github.com/wehw93/http-rest-api/internal/app/store/teststore"
)

func TestUserReposytoryCreate(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestFindUserByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err,store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
