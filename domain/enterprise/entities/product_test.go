package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: DISABLED,
		Price:  7240.30,
	}

	err := product.Enable()

	require.Nil(t, err)
	require.Equal(t, product.GetStatus(), ENABLED)

}

func TestProduct_Enable_Invalid_Price(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: DISABLED,
		Price:  0,
	}

	err := product.Enable()

	require.Error(t, err)
	require.Equal(t, "the price must be greather than zero to enable the product", err.Error())
	require.Equal(t, product.GetStatus(), DISABLED)

}

func TestProduct_Enable_Invalid_Price_Negative(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: DISABLED,
		Price:  -10,
	}

	err := product.Enable()

	require.Error(t, err)
	require.Equal(t, "the price must be greather than zero to enable the product", err.Error())
	require.Equal(t, product.GetStatus(), DISABLED)
}

func TestProduct_Disable(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: ENABLED,
		Price:  0,
	}

	err := product.Disable()

	require.Nil(t, err)
	require.Equal(t, product.GetStatus(), DISABLED)
}

func TestProduct_Disable_Invalid_Price(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: ENABLED,
		Price:  0.01,
	}

	err := product.Disable()

	require.Equal(t, err.Error(), "the price must be equals zero to disable the product")
	require.Equal(t, product.GetStatus(), ENABLED)
}

func TestProduct_Disable_Invalid_Price_Negative(t *testing.T) {
	product := Product{
		Name:   "iphone 15",
		Status: ENABLED,
		Price:  -0.001,
	}

	err := product.Disable()

	require.Equal(t, err.Error(), "the price must be equals zero to disable the product")
	require.Equal(t, product.GetStatus(), ENABLED)
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{
		Name:   "Xbox series x",
		Price:  3500,
		ID:     uuid.New().String(),
		Status: DISABLED,
	}

	result, err := product.IsValid()

	require.Nil(t, err)
	require.Equal(t, result, true)
}

func TestProduct_IsValid_InvalidID(t *testing.T) {
	product := Product{
		Name:   "Xbox series x",
		Price:  3500,
		ID:     "",
		Status: DISABLED,
	}

	result, err := product.IsValid()

	require.Equal(t, result, false)
	require.Equal(t, err.Error(), "ID: Missing required field")
}

func TestProduct_IsValid_InvalidName(t *testing.T) {
	product := Product{
		Name:   "",
		Price:  3500,
		ID:     uuid.NewString(),
		Status: DISABLED,
	}

	result, err := product.IsValid()

	require.Equal(t, result, false)
	require.Equal(t, err.Error(), "Name: non zero value required")
}

func TestProduct_IsValid_InvalidPrice(t *testing.T) {
	product := Product{
		Name:   "",
		Price:  -1,
		ID:     uuid.NewString(),
		Status: DISABLED,
	}

	result, err := product.IsValid()

	require.Equal(t, result, false)
	require.Equal(t, err.Error(), "the price must be greater or equal zero")
}

func TestProduct_IsValid_InvalidStatus(t *testing.T) {
	product := Product{
		Name:   "Xbox series x",
		Price:  3500,
		ID:     uuid.New().String(),
		Status: "invalid",
	}

	result, err := product.IsValid()

	require.Equal(t, result, false)
	require.Equal(t, err.Error(), "the status must be enabled or disabled")
}
