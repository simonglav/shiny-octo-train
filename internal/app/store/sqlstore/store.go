package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/simonglav/http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New - creates a new DB
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User - approaches to user repository interface, e.g. store.User().Create
// If it doesn't exist - creates new one
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}
