package entity

import (
	"github.com/google/uuid"
	"time"
)

type Installment struct {
	ID      uuid.UUID
	loadID  uuid.UUID
	amount  uint64
	paid    bool
	dueDate time.Time
}
