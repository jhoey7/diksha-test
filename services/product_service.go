package services

import (
	"edot-test/models"
	"edot-test/utils"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ProductProcessor interface {
	FindProductByQueryParam(req models.ProductListRequest) (product []models.Products, err error)
	CountProductByQueryParam(req models.ProductListRequest) (int, error)
	FindByPubID(pubID string) (product models.Products, err error)
	UpdateColumns(p models.Products, cols ...string) error
	FindByPubIDs(pubIDs []string) (product []models.Products, err error)
}

type ProductDetailProcessor interface {
	UpdateColumns(p models.ProductDetails, cols ...string) error
	FindByPubIds(ids []int) ([]models.ProductDetails, error)
}

// ProductService struct
type ProductService struct {
	Identifier       int64
	productProcessor ProductProcessor
}

// NewListProductService is func for initialize ProductService
func NewListProductService(pp ProductProcessor, i int64) ProductService {
	return ProductService{
		productProcessor: pp,
		Identifier:       i,
	}
}

// List is func for get list of product
func (svc ProductService) List(req models.ProductListRequest) models.Response {
	logs.Info("productListRequest : %+v", req)

	count, err := svc.productProcessor.CountProductByQueryParam(req)
	if err != nil {
		logs.Warn("[%d] Failed to find count of product: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	product, err := svc.productProcessor.FindProductByQueryParam(req)
	if err != nil && !errors.Is(err, orm.ErrNoRows) {
		logs.Warn("[%d] Failed to find product: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	var content models.ProductResponseWithTotal
	content.Data = product
	content.Total = count

	return models.ResponseSuccess(content)
}
