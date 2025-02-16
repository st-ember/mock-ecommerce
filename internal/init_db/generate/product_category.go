package generate

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db/model"
)

func ProductCategoryBatch() []*model.ProductCategory {
	categoryBatch := make([]*model.ProductCategory, 100)

	for i := 0; i < 100; i++ {
		category := &model.ProductCategory{
			Id: uuid.New(),
			Name: gofakeit.ProductCategory(),
			Slug: "",
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}
		categoryBatch[i] = category
	}

	return categoryBatch
}
