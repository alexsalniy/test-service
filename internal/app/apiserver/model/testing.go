package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestExtFIO(t *testing.T) *ExtendedFIO {
	t.Helper()

	return &ExtendedFIO{
		ID: uuid.New(),
		Name: "Name",
		Surname: "Surname",
	}
}