package models

import "time"

// Products struct for products
type Products struct {
	PubID         string            `orm:"pk;column(pubid)" json:"pubid"`
	Code          string            `orm:"column(product_code);null" json:"code"`
	Name          string            `orm:"column(product_name);null" json:"name"`
	Description   string            `orm:"column(product_description);null" json:"description"`
	TotalStock    int               `orm:"column(total_stock);null" json:"totalStock"`
	CreatedAt     time.Time         `orm:"column(created_at);null" json:"createdAt"`
	UpdatedAt     time.Time         `orm:"column(updated_at);null" json:"updatedAt"`
	ProductDetail []*ProductDetails `orm:"column(product_pubid);reverse(many);null" json:"details"`
}

// TableName for users
func (p *Products) TableName() string {
	return "products"
}

// ProductListRequest struct
type ProductListRequest struct {
	Sort    string
	Order   string
	Page    int
	Size    int
	Keyword string
}

// ProductResponseWithTotal response with total & list of data
type ProductResponseWithTotal struct {
	Total int        `json:"total"`
	Data  []Products `json:"data"`
}
