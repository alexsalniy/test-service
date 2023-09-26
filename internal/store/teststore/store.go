package teststore

import (
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/store"
)

type Store struct {
	extFIORepository *ExtFIORepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) ExtFIO() store.ExtendedFIORepository {
	if s.extFIORepository != nil {
		return s.extFIORepository
	}

	s.extFIORepository = &ExtFIORepository{
		store: s,
		fio: make(map[int]*model.ExtendedFIO),
	}

	return s.extFIORepository
}