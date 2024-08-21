package services

import (
	"edot-test/models"
	"edot-test/queue/publisher"
	"edot-test/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

// OrderProcessor interface
type OrderProcessor interface {
	Insert(o models.Orders) (models.Orders, error)
	FindPendingOrderByPubId(pubID string) (models.Orders, error)
	UpdateColumns(p models.Orders, cols ...string) error
}

// OrderDetailProcessor interface
type OrderDetailProcessor interface {
	InsertBulk(ods []models.OrdersDetail) ([]models.OrdersDetail, error)
}

// OrderService struct
type OrderService struct {
	Identifier             int64
	db                     orm.Ormer
	productProcessor       ProductProcessor
	productDetailProcessor ProductDetailProcessor
	orderProcessor         OrderProcessor
	orderDetailProcessor   OrderDetailProcessor
	releaseOrderQueue      func(message models.OrderReleaseQueue, delayTimeout int)
}

// NewCheckoutOrderService is func for initialize ProductService
func NewCheckoutOrderService(pp ProductProcessor, pdp ProductDetailProcessor, op OrderProcessor,
	odp OrderDetailProcessor, db orm.Ormer, i int64) OrderService {
	return OrderService{
		productProcessor:       pp,
		productDetailProcessor: pdp,
		orderProcessor:         op,
		orderDetailProcessor:   odp,
		db:                     db,
		Identifier:             i,
		releaseOrderQueue:      publisher.ReleaseOrderMessage,
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

	var productPubIDs []string
	mapProduct := make(map[string]models.CheckoutDetailRequest)
	for _, od := range request.Details {
		productPubIDs = append(productPubIDs, od.ProductPubID)
		mapProduct[od.ProductPubID] = od
	}

	products, err := svc.productProcessor.FindByPubIDs(productPubIDs)
	if err != nil {
		logs.Warn("[%d] Failed to find product pubIds: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	if len(request.Details) != len(products) {
		err = errors.New("product not found")
		logs.Warn("[%d] Failed to find product pubIds: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	for _, product := range products {
		totalStock := 0
		for _, detail := range product.ProductDetail {
			totalStock = totalStock + detail.Stock
		}

		if mapProduct[product.PubID].Qty <= 0 {
			logs.Warn("[%d] Quantity must be greater than zero for product: %s", svc.Identifier, product.Name)
			return models.ResponseError(fmt.Sprintf("Product %s out of stock.", product.Name), utils.ErrDefault)
		}

		if mapProduct[product.PubID].Qty > totalStock {
			logs.Warn("[%d] Insufficient stock for product: %s", svc.Identifier, product.PubID)
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

	var orderDtlReq []models.OrdersDetail
	remainingQty := make(map[string]int)
	for _, od := range request.Details {
		orderDtlReq = append(orderDtlReq, models.NewOrderDetail(order, od))
		remainingQty[od.ProductPubID] = od.Qty
	}

	_, err = svc.orderDetailProcessor.InsertBulk(orderDtlReq)
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to insert order detail: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	var releaseOrder []models.OrderRelease
	for _, product := range products {
		for _, detail := range product.ProductDetail {
			if remainingQty[detail.Product.PubID] == 0 {
				break
			}

			dtl := detail.NewProductDetail()
			stock := dtl.Stock

			qtyDeducted := 0
			if stock >= remainingQty[detail.Product.PubID] {
				dtl.Stock = stock - remainingQty[dtl.Product.PubID]
				err = svc.productDetailProcessor.UpdateColumns(dtl, "Stock")
				if err != nil {
					svc.db.Rollback()
					logs.Warn("[%d] Failed to update product: %s", svc.Identifier, err.Error())
					return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
				}

				qtyDeducted = remainingQty[dtl.Product.PubID]
				remainingQty[dtl.Product.PubID] = 0
			} else {
				dtl.Stock = 0
				err = svc.productDetailProcessor.UpdateColumns(dtl, "Stock")
				if err != nil {
					svc.db.Rollback()
					logs.Warn("[%d] Failed to update product: %s", svc.Identifier, err.Error())
					return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
				}

				qtyDeducted = stock
				remainingQty[dtl.Product.PubID] -= stock
			}

			releaseOrder = append(releaseOrder, models.OrderRelease{
				ProductDtlID: detail.ID,
				Qty:          qtyDeducted,
			})
		}
	}

	go svc.releaseOrderQueue(models.NewOrderReleaseQueue(order.PubID, releaseOrder), beego.AppConfig.DefaultInt("expOrder", 1))

	response := order

	svc.db.Commit()
	return models.ResponseSuccess(response)
}

func (svc OrderService) Confirm(pubID string) models.Response {
	order, err := svc.orderProcessor.FindPendingOrderByPubId(pubID)
	if err != nil {
		logs.Warn("[%d] Failed to find order: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	order.UpdatedAt = time.Now()
	order.Status = "CONFIRM"
	err = svc.orderProcessor.UpdateColumns(order, "Status", "UpdatedAt")
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to update order dtl: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	return models.ResponseSuccess(nil)
}

func (svc OrderService) ReleaseOrder(b []byte) models.Response {
	request := models.OrderReleaseQueue{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("release order request: %+v", request)

	order, err := svc.orderProcessor.FindPendingOrderByPubId(request.PubID)
	if err != nil {
		logs.Warn("[%d] Failed to find order: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	var ids []int
	mapsProductDtl := make(map[int]models.OrderRelease)
	for _, orq := range request.Data {
		ids = append(ids, orq.ProductDtlID)
		mapsProductDtl[orq.ProductDtlID] = orq
	}

	productDtlList, err := svc.productDetailProcessor.FindByPubIds(ids)
	if err != nil {
		logs.Warn("[%d] Failed to find product dtl: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	svc.db.Begin()
	for _, pd := range productDtlList {
		pd.Stock = pd.Stock + mapsProductDtl[pd.ID].Qty
		err := svc.productDetailProcessor.UpdateColumns(pd, "Stock")
		if err != nil {
			svc.db.Rollback()
			logs.Warn("[%d] Failed to update product dtl: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}
	}

	order.UpdatedAt = time.Now()
	order.Status = "CANCELED"
	err = svc.orderProcessor.UpdateColumns(order, "Status", "UpdatedAt")
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to update order dtl: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	svc.db.Commit()
	return models.ResponseSuccess(order)
}
