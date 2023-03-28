package aggregate

import (
	"azarm-lending-backend/entity"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidMember = errors.New("invalid member name")
)

type Member struct {
	entity.Person

	loans []*entity.Loan
}

func NewMember(name string) (Member, error) {
	if len(name) == 0 {
		return Member{}, ErrInvalidMember
	}

	member := Member{
		Person: entity.Person{
			ID:   uuid.New(),
			Name: name,
		},
		loans: make([]*entity.Loan, 0),
	}

	return member, nil
}
