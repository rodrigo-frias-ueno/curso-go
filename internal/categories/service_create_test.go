package categories_test

import (
	"context"
	"curso-go/internal/categories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
)

// CreateCategoryTestSuite defines the test suite
type CreateCategoryTestSuite struct {
	suite.Suite
	createRepoMock *MockCreateCategoryGateway
}

func (suite *CreateCategoryTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())

	suite.createRepoMock = NewMockCreateCategoryGateway(ctrl)
}

func (suite *CreateCategoryTestSuite) TestWithInvalidName() {
	ctx := context.TODO()
	candidate := categories.Category{}

	service := categories.New()
	_, err := service.CreateCategory(ctx, candidate)

	suite.Assert().NotEmpty(err)
	suite.Assert().EqualError(err, "invalid category: category name cannot be empty")
}

func (suite *CreateCategoryTestSuite) TestWithRepositoryError() {
	ctx := context.TODO()
	candidate := categories.Category{
		Name: "Test Category",
	}

	suite.createRepoMock.EXPECT().Create(ctx, gomock.Any()).Return(fmt.Errorf("Oops something went terrible wrong!!")).Times(1)

	service := categories.New(categories.WithCreateRepository(suite.createRepoMock))
	_, err := service.CreateCategory(ctx, candidate)

	suite.Assert().NotEmpty(err)
	suite.Assert().EqualError(err, "internal error while processing category")
}

func (suite *CreateCategoryTestSuite) TestSuccess() {
	ctx := context.TODO()
	candidate := categories.Category{
		Name: "Test Category",
	}

	suite.createRepoMock.EXPECT().Create(ctx, gomock.Any()).Return(nil).Times(1)

	service := categories.New(categories.WithCreateRepository(suite.createRepoMock))
	createdCategory, err := service.CreateCategory(ctx, candidate)

	suite.Assert().Empty(err)
	suite.Assert().Equal(createdCategory.Name, candidate.Name)
	suite.Assert().NotEmpty(createdCategory.ID)
	suite.Assert().Equal(createdCategory.Status, categories.StatusActive)
	suite.Assert().NotEmpty(createdCategory.CreatedAt)
	suite.Assert().NotEmpty(createdCategory.UpdatedAt)
	suite.Assert().Equal(createdCategory.CreatedAt, createdCategory.UpdatedAt)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CreateCategoryTestSuite))
}
