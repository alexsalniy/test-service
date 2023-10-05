package store

import (
	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
)

type ExtendedFIORepository interface {
	Create(*model.ExtendedFIO) error
	FindByID(*model.ExtendedFIO) error
}