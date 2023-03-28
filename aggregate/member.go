package aggregate

import "azarm-lending-backend/entity"

type Member struct {
	entity.Person

	loans []*entity.Loan
}
