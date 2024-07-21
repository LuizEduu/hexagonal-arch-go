package services

import (
	e "github.com/luizeduu/hexagonal-arch-go/domain/enterprise/entities"
	ps "github.com/luizeduu/hexagonal-arch-go/domain/enterprise/services"
)

type ProductService struct {
	Persitence ps.ProductPersitenceInterface
}

func (s *ProductService) Get(id string) (e.ProductInterface, error) {
	product, err := s.Persitence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Save(p e.ProductInterface) (e.ProductInterface, error) {
	_, err := p.IsValid()

	if err != nil {
		return nil, err
	}

	product, err := s.Persitence.Save(p)

	if err != nil {
		return nil, err
	}

	return product, nil
}
