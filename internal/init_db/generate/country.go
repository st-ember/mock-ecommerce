package generate

import (
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db/model"
)

func CountryBatch() []*model.Country{
	var countryBatch []*model.Country

	for i := 0; i < 100; i++ {
		country := &model.Country{
			Id: uuid.New(),
			Name: gofakeit.Country(),
			Code: gofakeit.Country(),
		}
		countryBatch = append(countryBatch, country)
	}
	return countryBatch
}
