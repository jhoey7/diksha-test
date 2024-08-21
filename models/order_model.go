package models

import (
	"edot-test/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

// Orders struct for orders
type Orders struct {
	PubID       string          `orm:"pk;column(pubid)" json:"pubid"`
	OrderNo     string          `orm:"column(order_no);null" json:"orderNo"`
	OrderDate   time.Time       `orm:"column(order_date);null" json:"orderDate"`
	OrderExp    time.Time       `orm:"column(order_expired);null" json:"orderExp"`
	UserPubID   string          `orm:"column(user_pubid);null" json:"userPubId"`
	Status      string          `orm:"column(status);null" json:"status"`
	CreatedAt   time.Time       `orm:"column(created_at);null" json:"createdAt"`
	UpdatedAt   time.Time       `orm:"column(updated_at);null" json:"updatedAt"`
	OrderDetail []*OrdersDetail `orm:"column(order_pubid);reverse(many);null" json:"details"`
}

// TableName for users
func (o *Orders) TableName() string {
	return "orders"
}

// NewPendingOrder is func for initialize Order
func NewPendingOrder(userPubID string) Orders {
	pubID := utils.NewV4().String()
	now := time.Now()
	orderExp := beego.AppConfig.DefaultInt("expOrder", 1)
	return Orders{
		PubID:     pubID,
		OrderNo:   fmt.Sprintf("%X", time.Now().UnixNano()),
		OrderDate: now,
		OrderExp:  now.Add(time.Minute * time.Duration(orderExp)),
		UserPubID: userPubID,
		Status:    "PENDING",
		CreatedAt: time.Now(),
	}
}

// CheckoutDetailRequest struct
type CheckoutDetailRequest struct {
	ProductPubID string `json:"productPubId"`
	Qty          int    `json:"qty"`
}

// CheckoutRequest struct
type CheckoutRequest struct {
	UserPubID string                  `json:"userPubId"`
	Details   []CheckoutDetailRequest `json:"product"`
}
