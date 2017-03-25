package country

import (
	"errors"
	"fmt"
)

// CountryService represents the feature: Inserting Number
type CountryService interface {
	Visit(string) (string, error)
}

// service represents CountryService Instance
type service struct {
	repository CountryRepository
}

func NewCountryService(r CountryRepository) CountryService {
	return &service{repository: r}
}

func (svc *service) Visit(country string) (string, error) {
	if country == "" {
		return "", errors.New("Empty `user`")
	}

	count, err := svc.repository.Visit(country)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s has been visited by %d people", country, count), nil
}
