package entity

import (
	"github.com/google/uuid"
	"time"
)

type Loan struct {
	ID               uuid.UUID
	amount           uint64
	noOfInstallments int16
	createdAt        time.Time
}
