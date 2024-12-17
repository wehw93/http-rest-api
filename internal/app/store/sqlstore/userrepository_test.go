package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store"
	"github.com/wehw93/http-rest-api/internal/app/store/sqlstore"
)

func TestUserReposytoryCreate(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, database_url)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestFindUserByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, database_url)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
