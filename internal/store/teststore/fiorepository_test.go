package teststore_test

import (
	"testing"

	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/store"
	"github.com/alexsalniy/test-service/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestFIORepository_Create(t *testing.T) {
	s := teststore.New()
	e := model.TestExtFIO(t)
	assert.NoError(t, s.ExtFIO().Create(e))
	assert.NotNil(t, e.ID)
}

func TestFIORepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	e := model.TestExtFIO(t)
	err := s.ExtFIO().FindByID(e)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.ExtFIO().Create(e)
	err = s.ExtFIO().FindByID(e)
	assert.NoError(t, err)
	assert.NotNil(t, e)
}