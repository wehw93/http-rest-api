package teststore

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ..
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		Store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
