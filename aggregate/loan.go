package aggregate

import (
	"azarm-lending-backend/entity"
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidLoan = errors.New("invalid loan")
)

type Loan struct {
	entity.Loan

	installments []*entity.Installment
}

func NewLoan(amount uint64, noOfInstallments uint16) (Loan, error) {
	if amount == 0 || noOfInstallments == 0 {
		return Loan{}, ErrInvalidLoan
	}

	loan := Loan{
		Loan: entity.Loan{
			ID:               uuid.New(),
			Amount:           amount,
			NoOfInstallments: noOfInstallments,
			CreatedAt:        time.Now(),
		},
		installments: make([]*entity.Installment, 0),
	}

	return loan, nil
}
