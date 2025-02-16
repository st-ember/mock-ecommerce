package storage

import (
	"github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/init_db/generate"
)

func StoreInitCountries() error {
	countryBatch := generate.CountryBatch()
	idArr :=  make([]uuid.UUID, len(countryBatch))
	nameArr := make([]string, len(countryBatch)) 
	codeArr := make([]string, len(countryBatch))

	for i, country := range countryBatch {
		idArr[i] = country.Id
		nameArr[i] = country.Name
		codeArr[i] = country.Code
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO country (id, name, code)
	SELECT * FROM UNNEST($1::uuid[], $2::text[], $3::text[])`

	_, err = tx.Exec(query, pq.Array(idArr), pq.Array(nameArr), pq.Array(codeArr))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func CountryIds() ([]uuid.UUID, error) {
	var countryIds []uuid.UUID

	query := "SELECT id FROM country"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		countryIds = append(countryIds, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return countryIds, nil
}
