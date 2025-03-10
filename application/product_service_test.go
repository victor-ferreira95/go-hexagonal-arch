package application_test

import (
	"go-hexagonal/application"
	mock_application "go-hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	// Arrange
	// Act
	// Assert
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
