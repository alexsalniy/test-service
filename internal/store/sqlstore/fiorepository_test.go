package sqlstore_test

import (
	"testing"

	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/store/sqlstore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestingDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	e := model.TestExtFIO(t)

	assert.NoError(t, s.ExtFIO().Create(e))
	assert.NotNil(t, e)
}

func TestFIORepository_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestingDB(t, databaseURL)
	defer teardown("fio")

	id := uuid.New()

	s := sqlstore.New(db)
	_, err := s.ExtFIO().FindByID(id)
	assert.Error(t, err)

	s.ExtFIO().Create(&model.ExtendedFIO{
		ID: uuid.New(),
		Name: "Name",
	})

	e, err := s.ExtFIO().FindByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)


}