package services

import (
	"testing"

	application_mocks "github.com/luizeduu/hexagonal-arch-go/test/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductServiceGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := application_mocks.NewMockProductInterface(ctrl)
	persistence := application_mocks.NewMockProductPersitenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := ProductService{
		Persitence: persistence,
	}

	result, err := service.Get("any_id")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := application_mocks.NewMockProductInterface(ctrl)
	persistence := application_mocks.NewMockProductPersitenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := ProductService{
		Persitence: persistence,
	}

	result, err := service.Create("product name", 4201.22)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServiceEnable(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := application_mocks.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)

	persistence := application_mocks.NewMockProductPersitenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := ProductService{
		Persitence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServiceDisable(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := application_mocks.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)

	persistence := application_mocks.NewMockProductPersitenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := ProductService{
		Persitence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}
