package models

// OrdersDetail struct for order_detail
type OrdersDetail struct {
	ID           int     `orm:"auto;pk;column(id)" json:"id"`
	Order        *Orders `orm:"column(order_pubid);rel(fk)" json:"-"`
	ProductPubID string  `orm:"column(product_pubid)" json:"productPubId"`
	Qty          int     `orm:"column(qty);null" json:"qty"`
}

// TableName for users
func (od *OrdersDetail) TableName() string {
	return "orders_details"
}

// NewOrderDetail is func for initialize OrderDetail
func NewOrderDetail(order Orders, cdr CheckoutDetailRequest) OrdersDetail {
	return OrdersDetail{
		Order:        &order,
		ProductPubID: cdr.ProductPubID,
		Qty:          cdr.Qty,
	}
}
