package number

import (
	"fmt"
	"sync"

	. "github.com/a-trium/gipeline/server-gateway/service/common"
)

type NumberRepository interface {
	Store(s *Submission) error
	Find(u User) (*Submission, error)
	FindAll() []*Submission
}

type numberRepositoryInst struct {
	mtx         sync.RWMutex
	submissions map[User]*Submission
}

func NewNumberRepository() NumberRepository {
	return &numberRepositoryInst{
		submissions: make(map[User]*Submission),
	}
}

func (r *numberRepositoryInst) Store(s *Submission) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	subs, exist := r.submissions[s.User]
	if !exist {
		r.submissions[s.User] = s
	} else {
		subs.Update(s.Number)
	}

	return nil
}

func (r *numberRepositoryInst) Find(u User) (*Submission, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	s, exist := r.submissions[u]
	if !exist {
		return nil, fmt.Errorf("Can't find User: %s", u)
	}

	return s, nil
}

func (r *numberRepositoryInst) FindAll() []*Submission {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	ss := make([]*Submission, 0, len(r.submissions))
	for _, s := range r.submissions {
		ss = append(ss, s)
	}

	return ss
}
