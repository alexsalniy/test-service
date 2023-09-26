package teststore

import (
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/google/uuid"
)

type ExtFIORepository struct {
	store *Store
	fio   map[int]*model.ExtendedFIO
}

// FindByID implements store.ExtendedFIORepository.
func (*ExtFIORepository) FindByID(uuid.UUID) (*model.ExtendedFIO, error) {
	panic("unimplemented")
}

func (r *ExtFIORepository) Create(e *model.ExtendedFIO) error {

	return nil
}

