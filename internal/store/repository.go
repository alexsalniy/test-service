package store

import (
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/google/uuid"
)

type ExtendedFIORepository interface {
	Create(*model.ExtendedFIO) error
	FindByID(uuid.UUID) (*model.ExtendedFIO, error)
}