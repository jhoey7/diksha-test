package services

import (
	"diskha-test/models"
	"diskha-test/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// OrderProcessor interface
type OrderProcessor interface {
	GetTotalQtyPendingOrderByProductPubID(pID string) (count int, err error)
	Insert(o models.Orders) (models.Orders, error)
}

// OrderDetailProcessor interface
type OrderDetailProcessor interface {
	Insert(od models.OrdersDetail) (models.OrdersDetail, error)
}

// OrderService struct
type OrderService struct {
	Identifier           int64
	db                   orm.Ormer
	productProcessor     ProductProcessor
	orderProcessor       OrderProcessor
	orderDetailProcessor OrderDetailProcessor
}

// NewCheckoutOrderService is func for initialize ProductService
func NewCheckoutOrderService(pp ProductProcessor, op OrderProcessor,
	odp OrderDetailProcessor, db orm.Ormer, i int64) OrderService {
	return OrderService{
		productProcessor:     pp,
		orderProcessor:       op,
		orderDetailProcessor: odp,
		db:                   db,
		Identifier:           i,
	}
}

// Checkout is func for checkout order
func (svc OrderService) Checkout(b []byte) models.Response {
	request := models.CheckoutRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("checkout request: %+v", request)

	stock := 0
	for _, d := range request.Details {
		count, err := svc.orderProcessor.GetTotalQtyPendingOrderByProductPubID(d.ProductPubID)
		if err != nil {
			logs.Warn("[%d] Failed to find count of pending order: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		product, err := svc.productProcessor.FindByPubID(d.ProductPubID)
		if err != nil {
			logs.Warn("[%d] Failed to find product: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		stock = product.TotalStock - count
		if d.Qty > stock {
			logs.Warn("[%d] Out of stock product: %s", svc.Identifier, d.ProductPubID)
			return models.ResponseError(fmt.Sprintf("Product %s out of stock.", product.Name), utils.ErrDefault)
		}
	}

	svc.db.Begin()
	orderReq := models.NewPendingOrder(request.UserPubID)
	order, err := svc.orderProcessor.Insert(orderReq)
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to insert order: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	for _, od := range request.Details {
		dtlOrderReq := models.NewOrderDetail(order, od)
		_, err = svc.orderDetailProcessor.Insert(dtlOrderReq)
		if err != nil {
			svc.db.Rollback()
			logs.Warn("[%d] Failed to insert order detail: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		product, err := svc.productProcessor.FindByPubID(od.ProductPubID)
		if err != nil {
			svc.db.Rollback()
			logs.Warn("[%d] Failed to find product: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		product.TotalStock = product.TotalStock - od.Qty
		err = svc.productProcessor.UpdateColumns(product, "TotalStock")
		if err != nil {
			svc.db.Rollback()
			logs.Warn("[%d] Failed to update product: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}
	}

	svc.db.Commit()
	return models.ResponseSuccess(nil)
}
