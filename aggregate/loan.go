package aggregate

import "azarm-lending-backend/entity"

type Loan struct {
	entity.Loan

	installments []*entity.Installment
}
