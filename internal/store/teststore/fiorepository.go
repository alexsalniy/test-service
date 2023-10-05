package teststore

import (
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
)

type ExtFIORepository struct {
	store *Store
	fio   map[int]*model.ExtendedFIO
}

func (*ExtFIORepository) FindByID(e *model.ExtendedFIO) error {
	panic("unimplemented")
}

func (r *ExtFIORepository) Create(e *model.ExtendedFIO) error {

	return nil
}

