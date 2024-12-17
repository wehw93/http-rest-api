package teststore

import (
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store"
)

type UserRepository struct {
	Store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil

}
