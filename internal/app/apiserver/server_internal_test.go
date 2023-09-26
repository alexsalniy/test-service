package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexsalniy/test-service/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/fio", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}