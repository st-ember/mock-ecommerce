package generate

import (
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db/model"
)
 
func CustomerBatch() ([]*model.Customer, error) {
	customerBatch := make([]*model.Customer, 1000)
	
	for i := 0; i < 1000; i++ {
		customer := &model.Customer{
			Id: uuid.New(),
			Username: gofakeit.Username(),
			JoinedAt: gofakeit.Date(),
		}
		customerBatch[i] = customer
	}
	return customerBatch, nil
}