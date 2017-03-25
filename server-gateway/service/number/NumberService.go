package number

import (
	"errors"
	"fmt"

	. "github.com/a-trium/gipeline/server-gateway/service/common"
)

// NumberService represents the feature: Inserting Number
type NumberService interface {
	Insert(string, int) (string, error)
	Total() int
}

// service represents NumberService Instance
type service struct {
	repository NumberRepository
}

func NewNumberService(r NumberRepository) NumberService {
	return &service{repository: r}
}

func (svc *service) Insert(u string, n int) (string, error) {
	if u == "" {
		return "", errors.New("Empty `user`")
	}

	user := User(u)
	number := Number(n)
	subs := &Submission{User: user, Number: number}

	svc.repository.Store(subs)
	s, err := svc.repository.Find(user)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s has been inserted: %d", user, s.Number), nil
}

func (svc *service) Total() int {
	var total int = 0

	subs := svc.repository.FindAll()

	for _, s := range subs {
		total += s.GetNumber()
	}

	return total
}
