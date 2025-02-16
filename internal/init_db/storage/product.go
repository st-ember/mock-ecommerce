package storage

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/init_db/generate"
)

func StoreInitProducts() error {
	productBatch, err := generate.ProductBatch()
	if err != nil {
		return err
	}

	categoryIds, err := ProductCategoryIds()
	if err != nil {
		return err
	}

	merchantIds, err := MerchantIds()
	if err != nil {
		return err
	}
	
	// create arrays for unnest
	uuidArr := make([]uuid.UUID, len(productBatch))
	categoryArr := make([]uuid.UUID, len(productBatch))
	nameArr := make([]string, len(productBatch))
	descArr := make([]string, len(productBatch))
	merchantArr := make([]uuid.UUID, len(productBatch))
	priceArr := make([]float32, len(productBatch))
	createDateArr := make([]time.Time, len(productBatch))
	updateDateArr := make([]time.Time, len(productBatch))
	
	for i, product  := range productBatch {
		randCatIdx := rand.Intn(len(categoryIds))
		randMerchIdx := rand.Intn(len(merchantIds))
		
		uuidArr[i] = product.Id
		categoryArr[i] = categoryIds[randCatIdx]
		nameArr[i] = product.Name
		descArr[i] = product.Description
		merchantArr[i] = merchantIds[randMerchIdx]
		priceArr[i] = product.Price
		createDateArr[i] = product.CreatedAt
		updateDateArr[i] = product.UpdatedAt
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
	
	query := `INSERT INTO product (id, category, name, description, merchant, price, created_at, updated_at)
	SELECT * FROM UNNEST($1::uuid[], $2::uuid[], $3::text[], $4::text[], $5::uuid[], $6::float[], $7::timestamp[], $8::timestamp[])`
	
	_, err = tx.Exec(query, pq.Array(uuidArr), pq.Array(categoryArr), 
		pq.Array(nameArr), pq.Array(descArr), pq.Array(merchantArr), pq.Array(priceArr), 
		pq.Array(createDateArr), pq.Array(updateDateArr))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
