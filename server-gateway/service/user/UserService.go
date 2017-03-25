package user

import (
	"errors"

	. "github.com/a-trium/gipeline/server-gateway/service/common"
	"github.com/a-trium/gipeline/server-gateway/service/number"
)

type UserService interface {
	Users() []User
	User(string) (int, error)
}

type service struct {
	repository number.NumberRepository
}

func NewUserService(r number.NumberRepository) UserService {
	return &service{repository: r}
}

func (svc *service) Users() []User {
	var users []User

	for _, subs := range svc.repository.FindAll() {
		users = append(users, subs.User)
	}

	return users
}

func (svc service) User(u string) (int, error) {
	if u == "" {
		return 0, errors.New("Empty `user`")
	}

	user := User(u)
	subs, err := svc.repository.Find(user)
	if err != nil {
		return 0, err
	}

	return subs.GetNumber(), err
}
