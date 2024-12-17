package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/wehw93/http-rest-api/internal/app/store/sqlstore"
)

func Start(config *Config) error {

	db, err := NewDB(config.dataBaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)

}
func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
