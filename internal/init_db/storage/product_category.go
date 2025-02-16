package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/init_db/generate"
)

func StoreInitProductCategories() error {
	productCategoryBatch := generate.ProductCategoryBatch()

	// create arrays for unnest
	uuidArr := make([]uuid.UUID, len(productCategoryBatch))
	nameArr := make([]string, len(productCategoryBatch))
	slugArr := make([]string, len(productCategoryBatch))
	createDateArr := make([]time.Time, len(productCategoryBatch))
	updateDateArr := make([]time.Time, len(productCategoryBatch))
	
	for i, category  := range productCategoryBatch {
		uuidArr[i] = category.Id
		nameArr[i] = category.Name
		slugArr[i] = category.Slug
		createDateArr[i] = category.CreatedAt
		updateDateArr[i] = category.UpdatedAt
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
	
	query := `INSERT INTO product_category (id, name, slug, created_at, updated_at)
	SELECT * FROM UNNEST($1::uuid[], $2::text[], $3::text[], $4::timestamp[], $5::timestamp[])`
	
	_, err = tx.Exec(query, pq.Array(uuidArr), pq.Array(nameArr), pq.Array(slugArr), pq.Array(createDateArr), pq.Array(updateDateArr))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func ProductCategoryIds() ([]uuid.UUID, error) {
	var productCategoryIds []uuid.UUID

	query := "SELECT id FROM product_category"

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

		productCategoryIds = append(productCategoryIds, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productCategoryIds, nil
}
