package sqlstore

import (
	"log"

	"github.com/alexsalniy/test-service/internal/app/apiserver/model"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type ExtFIORepository struct {
	store *Store
}

func (r *ExtFIORepository) Create(e *model.ExtendedFIO) error {
	if e.Validator() {


		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	
		e.GetExtension("AGEAPI")
		e.GetExtension("GENDERAPI")
		e.GetExtension("NATIONAPI")
	
		e.ID = uuid.New()
	
		if err := r.store.db.QueryRow(`
			INSERT INTO fio (id, name, surname, patronymic, age, gender, gender_probability)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id`,
			e.ID, 
			e.Name,
			e.Surname,
			e.Patronymic,
			e.Age,
			e.Gender,
			e.Probability,
		).Scan(&e.ID); err != nil {
			return err
		}
	
		for i := 0; i < len(e.Nation); i++ {
	
			if err := r.store.db.QueryRow(`
				ALTER TABLE fio ADD COLUMN IF NOT EXISTS %s NUMERIC`, 
				e.Nation[i].CountryID,
			).Scan(); err != nil {
				return err
			}
				
			if err := r.store.db.QueryRow(`
				UPDATE fio 
				SET %s = %g
				WHERE id = '%v'`, 
				e.Nation[i].CountryID, 
				e.Nation[i].Probability, 
				e.ID,
			).Scan(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ExtFIORepository) FindByID(e *model.ExtendedFIO) error {
	if err := r.store.db.QueryRow(`
	SELECT id, name, surname, patronymic, age, gender, gender_probability
	FROM fio
	WHERE id = $1`, e.ID,
).Scan(
	&e.ID, 
	&e.Name,
	&e.Surname,
	&e.Patronymic,
	&e.Age,
	&e.Gender,
	&e.Probability,
); err != nil {
	return err
}
	return nil
}