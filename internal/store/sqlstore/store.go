package sqlstore

import (
	"database/sql"

	"github.com/alexsalniy/test-service/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db 								*sql.DB
	extFIORepository 	*ExtFIORepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) ExtFIO() store.ExtendedFIORepository {
	if s.extFIORepository != nil {
		return s.extFIORepository
	}

	s.extFIORepository = &ExtFIORepository{
		store: s,
	}

	return s.extFIORepository
}
