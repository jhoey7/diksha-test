package models

import (
	"diskha-test/utils"
	"fmt"
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
	OrderDetail []*OrdersDetail `orm:"column(order_pubid);reverse(many);null" json:"details"`
}

// TableName for users
func (o *Orders) TableName() string {
	return "orders"
}

// NewPendingOrder is func for initialize Order
func NewPendingOrder(userPubID string) Orders {
	pubID := utils.NewV4().String()
	return Orders{
		PubID:     pubID,
		OrderNo:   fmt.Sprintf("%X", time.Now().UnixNano()),
		OrderDate: time.Now(),
		OrderExp:  time.Time{},
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
