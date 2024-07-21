package services

import e "github.com/luizeduu/hexagonal-arch-go/domain/enterprise/entities"

type ProductServiceInterface interface {
	Get(id string) (e.ProductInterface, error)
	Create(product e.ProductInterface) (e.ProductInterface, error)
	Enable(product e.ProductInterface) (e.ProductInterface, error)
	Disable(product e.ProductInterface) (e.ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (e.ProductInterface, error)
}

type ProductWriter interface {
	Save(product e.ProductInterface) (e.ProductInterface, error)
}

type ProductPersitenceInterface interface {
	ProductReader
	ProductWriter
}
