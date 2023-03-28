package entity

import (
	"github.com/google/uuid"
	"time"
)

type Installment struct {
	ID      uuid.UUID
	LoanID  uuid.UUID
	Amount  uint64
	Paid    bool
	DueDate time.Time
	PayDate time.Time
}
