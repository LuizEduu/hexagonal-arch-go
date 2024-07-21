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

func (s *ProductService) Create(name string, price float64) (e.ProductInterface, error) {
	product := e.NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()

	if err != nil {
		return &e.Product{}, err
	}

	result, err := s.Persitence.Save(product)

	if err != nil {
		return &e.Product{}, err
	}

	return result, nil
}

func (s *ProductService) Enable(product e.ProductInterface) (e.ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return &e.Product{}, err
	}

	result, err := s.Persitence.Save(product)

	if err != nil {
		return &e.Product{}, err
	}

	return result, nil
}

func (s *ProductService) Disable(product e.ProductInterface) (e.ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return &e.Product{}, err
	}

	result, err := s.Persitence.Save(product)

	if err != nil {
		return &e.Product{}, err
	}

	return result, nil
}
