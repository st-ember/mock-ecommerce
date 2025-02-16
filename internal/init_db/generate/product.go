package generate

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db/model"
)
 
func ProductBatch() ([]*model.Product, error) {
	customerBatch := make([]*model.Product, 1000)
	
	for i := 0; i < 1000; i++ {
		customer := &model.Product{
			Id: uuid.New(),
			Name: gofakeit.ProductName(),
			Description: gofakeit.Paragraph(1, 2, 10, "\n"),
			Price: float32(gofakeit.Price(1, 9999999999)),
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}
		customerBatch[i] = customer
	}
	return customerBatch, nil
}
