package storage

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/init_db/generate"
)

func StoreInitCustomers() error {
	// retrieve 1000 customers 
	customerBatch, err := generate.CustomerBatch()
	if err != nil {
		return err
	}

	countryIds, err := CountryIds()
	if err != nil {
		return err
	}
	
	// create arrays for unnest
	uuidArr := make([]uuid.UUID, len(customerBatch))
	usernameArr := make([]string, len(customerBatch))
	countryArr := make([]uuid.UUID, len(customerBatch))
	joinDateArr := make([]time.Time, len(customerBatch))
	
	for i, customer  := range customerBatch {
		randIdx := rand.Intn(len(countryIds))
		
		uuidArr[i] = customer.Id
		usernameArr[i] = customer.Username
		// assign random country ids to customerBatch
		countryArr[i] = countryIds[randIdx]
		joinDateArr[i] = customer.JoinedAt
	}
	
	// batch store into db
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	
	query := `INSERT INTO customer (id, username, country, joined_at)
	SELECT * FROM UNNEST($1::uuid[], $2::text[], $3::uuid[], $4::timestamp[])`
	
	_, err = tx.Exec(query, pq.Array(uuidArr), pq.Array(usernameArr), pq.Array(countryArr), pq.Array(joinDateArr))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
