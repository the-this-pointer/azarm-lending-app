package entity

import (
	"github.com/google/uuid"
	"time"
)

type Loan struct {
	ID               uuid.UUID
	Amount           uint64
	NoOfInstallments uint16
	CreatedAt        time.Time
}
