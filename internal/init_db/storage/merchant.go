package storage

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/db/model"
	"github.com/st-ember/mockecommerce/internal/init_db/generate"
)

func StoreInitMerchants() error {
	// retrieve 1000 merchants
	merchantBatch, err := generate.MerchantBatch()
	if err != nil {
		return err
	}

	countryIds, err := CountryIds()
	if err != nil {
		return err
	}
	
	// create arrays for unnest
	uuidArr := make([]uuid.UUID, len(merchantBatch))
	usernameArr := make([]string, len(merchantBatch))
	countryArr := make([]uuid.UUID, len(merchantBatch))
	descArr := make([]string, len(merchantBatch))
	statusArr := make([]model.MerchantStatus, len(merchantBatch))
	joinDateArr := make([]time.Time, len(merchantBatch))
	
	for i, merchant  := range merchantBatch {
		randIdx := rand.Intn(len(countryIds))
		
		uuidArr[i] = merchant.Id
		usernameArr[i] = merchant.Username
		// assign random country ids to merchantBatch
		countryArr[i] = countryIds[randIdx]
		descArr[i] = merchant.Description
		statusArr[i] = merchant.Status
		joinDateArr[i] = merchant.JoinedAt
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
	
	query := `INSERT INTO merchant (id, username, country, description, status, joined_at)
	SELECT * FROM UNNEST($1::uuid[], $2::text[], $3::uuid[], $4::text[], $5::int[], $6::timestamp[])`
	
	_, err = tx.Exec(query, pq.Array(uuidArr), pq.Array(usernameArr), pq.Array(countryArr), pq.Array(descArr), pq.Array(statusArr), pq.Array((joinDateArr)))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func MerchantIds() ([]uuid.UUID, error) {
	var ids []uuid.UUID

	query := "SELECT id FROM merchant"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}