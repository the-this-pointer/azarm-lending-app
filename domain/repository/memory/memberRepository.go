package memory

import (
	"azarm-lending-backend/aggregate"
	"azarm-lending-backend/domain/repository"
	"azarm-lending-backend/entity"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemberMemRepository struct {
	members map[uuid.UUID]aggregate.Member
	sync.Mutex
}

func New() *MemberMemRepository {
	return &MemberMemRepository{
		members: make(map[uuid.UUID]aggregate.Member),
	}
}

func (mr *MemberMemRepository) FindAll() ([]entity.Person, error) {
	mr.Lock()
	defer mr.Unlock()

	members := make([]entity.Person, len(mr.members))
	for _, v := range mr.members {
		members = append(members, v.Person)
	}
	return members, nil
}

func (mr *MemberMemRepository) Get(id uuid.UUID) (aggregate.Member, error) {
	if member, ok := mr.members[id]; ok {
		return member, nil
	}

	return aggregate.Member{}, repository.ErrMemberNotFound
}

func (mr *MemberMemRepository) Add(c aggregate.Member) error {
	if mr.members == nil {
		mr.Lock()
		mr.members = make(map[uuid.UUID]aggregate.Member)
		mr.Unlock()
	}
	if _, ok := mr.members[c.ID]; ok {
		return fmt.Errorf("member already exists: %w", repository.ErrFailedToAddMember)
	}
	mr.Lock()
	mr.members[c.ID] = c
	mr.Unlock()
	return nil
}

func (mr *MemberMemRepository) Update(c aggregate.Member) error {
	if _, ok := mr.members[c.ID]; !ok {
		return fmt.Errorf("member does not exist: %w", repository.ErrUpdateMember)
	}
	mr.Lock()
	mr.members[c.ID] = c
	mr.Unlock()
	return nil
}
