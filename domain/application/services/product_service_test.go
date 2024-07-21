package services

import (
	"testing"

	application_mocks "github.com/luizeduu/hexagonal-arch-go/domain/application/mocks"
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
