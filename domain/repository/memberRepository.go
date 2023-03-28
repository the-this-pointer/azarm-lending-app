package repository

import (
	"azarm-lending-backend/aggregate"
	"azarm-lending-backend/entity"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrMemberNotFound    = errors.New("the member was not found in the repository")
	ErrFailedToAddMember = errors.New("failed to add the member to the repository")
	ErrUpdateMember      = errors.New("failed to update the member in the repository")
)

type MemberRepository interface {
	FindAll() ([]entity.Person, error)
	Get(uuid uuid.UUID) (aggregate.Member, error)
	Add(member aggregate.Member) error
	Update(member aggregate.Member) error
}
