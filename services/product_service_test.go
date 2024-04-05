package services

import (
	"diskha-test/models"
	"diskha-test/utils"
	"errors"
	"testing"
)

type fakeProductProcessor struct {
	products              []models.Products
	errFindProduct        error
	countProduct          int
	errCountFindProduct   error
	product               models.Products
	errFindProductByPubID error
	errUpdate             error
}

func (f fakeProductProcessor) FindProductByQueryParam(req models.ProductListRequest) (products []models.Products, err error) {
	return f.products, f.errFindProduct
}

func (f fakeProductProcessor) CountProductByQueryParam(req models.ProductListRequest) (int, error) {
	return f.countProduct, f.errCountFindProduct
}

func (f fakeProductProcessor) FindByPubID(pubID string) (product models.Products, err error) {
	return f.product, f.errFindProductByPubID
}

func (f fakeProductProcessor) UpdateColumns(p models.Products, cols ...string) error {
	return f.errUpdate
}

var (
	positiveFindProduct        = fakeProductProcessor{}
	negativeFindCountOfProduct = fakeProductProcessor{
		errCountFindProduct: errors.New("DB Error"),
	}
	negativeFindProduct = fakeProductProcessor{
		errFindProduct: errors.New("DB Error"),
	}
)

func TestProductService_List(t *testing.T) {
	cases := []struct {
		testName         string
		productProcessor ProductProcessor
		expectedResponse int
	}{
		{
			testName:         "Positive Case: Success Flow",
			expectedResponse: utils.ErrNone,
			productProcessor: positiveFindProduct,
		},
		{
			testName:         "Negative Case: Failed find count of product",
			expectedResponse: utils.ErrDefault,
			productProcessor: negativeFindCountOfProduct,
		},
		{
			testName:         "Negative Case: Failed find product",
			expectedResponse: utils.ErrDefault,
			productProcessor: negativeFindProduct,
		},
	}

	for _, c := range cases {
		svc := NewListProductService(c.productProcessor, 123)
		resp := svc.List(models.ProductListRequest{})
		if resp.Code != c.expectedResponse {
			t.Errorf(expectedResponseCode, c.expectedResponse, resp.Code)
		}
	}
}
