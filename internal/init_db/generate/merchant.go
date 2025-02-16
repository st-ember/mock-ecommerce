package generate

import (
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/st-ember/mockecommerce/internal/db/model"
)
 
func MerchantBatch() ([]*model.Merchant, error) {
	merchantBatch := make([]*model.Merchant, 1000)
	
	for i := 0; i < 1000; i++ {
		merchant := &model.Merchant{
			Id: uuid.New(),
			Username: gofakeit.Username(),
			JoinedAt: gofakeit.Date(),
			Description: gofakeit.Paragraph(3, 8, 10, "\n"),
			Status: model.Active,
		}
		merchantBatch[i] = merchant
	}
	return merchantBatch, nil
}
