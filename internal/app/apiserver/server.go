package apiserver

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/alexsalniy/test-service/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

func newServer(store store.Store) *server {
	s :=  &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:	store,
	}

	var wg sync.WaitGroup
	Producer()

	wg.Add(2)
	go s.configureRouter(&wg)
	go Consumer(store, &wg)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter(wg *sync.WaitGroup) {
	defer wg.Done()
	s.router.HandleFunc("/fio", s.handleFIOCreate()).Methods("POST")
	s.router.HandleFunc("/id", s.handleFIOReturn()).Methods("GET")
}

func (s *server) handleFIOCreate() http.HandlerFunc {
	type request struct {
		Name				string `json:"name`
		Surname			string `json:"surname`
		Patronymic	string `json:"patronymic`
	}
	
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.ExtendedFIO{
			Name: req.Name,
			Surname: req.Surname,
			Patronymic: req.Patronymic,
		}

		if err := s.store.ExtFIO().Create(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		
		s.respond(w, r, http.StatusCreated, e)
	}
}

func (s *server) handleFIOReturn() http.HandlerFunc {
	type request struct {
		ID				uuid.UUID `json:"id`
	}
	
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.ExtendedFIO{
			ID: req.ID,
		}

		if err := s.store.ExtFIO().FindByID(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		
		s.respond(w, r, http.StatusCreated, e)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}