package country

import (
	"sync"
)

type CountryRepository interface {
	Visit(country string) (int, error)
}

type countryRepositoryInst struct {
	mtx         sync.RWMutex
	visitCount map[string]int
}

func NewCountryRepository() CountryRepository {
	return &countryRepositoryInst{
		visitCount: make(map[string]int),
	}
}

func (r *countryRepositoryInst) Visit(country string) (int, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	count, exist := r.visitCount[country]
	if !exist {
		count = 1
	} else {
		count = count + 1
	}
	r.visitCount[country] = count

	return count, nil
}

