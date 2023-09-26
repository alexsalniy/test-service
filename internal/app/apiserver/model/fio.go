package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Nation struct {
	CountryID 	string `json:"country_id"`
	Probability float32
}

type fioResponse struct {
	Count 			int 			`json:"count"`
	Name 				string 		`json:"name"`
	Age 				int				`json:"age"`
	Gender 			string		`json:"gender"`
	Probability float32		`json:"probability"`
	Nation 			[]Nation 	`json:"nation"`
}

type FIO struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type ExtendedFIO struct {
	ID 					uuid.UUID
	Name 				string 		`json:"name"`
	Surname 		string 		
	Patronymic 	string 
	Age 				int				`json:"age"`
	Gender 			string		`json:"gender"`
	Probability float32		`json:"probability"`
	Nation 			[]Nation 	`json:"nation"`
}

func (e *ExtendedFIO) Validator() bool {
  nameField := fmt.Sprint(reflect.TypeOf(e.Name)) 
  surnameField := fmt.Sprint(reflect.TypeOf(e.Surname)) 

  if nameField != "string" {
		log.Fatal("struct does not have the field '%s'.\n", e.Surname)
		return false
	}

  if surnameField != "string" {
		log.Fatal("struct does not have the field '%s'.\n", e.Surname)
		return false
  }
	return true
}

func (r *ExtendedFIO)GetExtension (es string) (*ExtendedFIO, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	res, err := http.Get(fmt.Sprintf(os.Getenv(es), r.Name))
	if err != nil {
			log.Fatal("Error making http request: %s\n", err)
	}

	rjson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	err = json.Unmarshal(rjson, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}