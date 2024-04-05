package models

import "time"

// ProductDetails struct for product_details
type ProductDetails struct {
	ID          int       `orm:"auto;pk;column(id)" json:"id"`
	Product     *Products `orm:"column(product_pubid);rel(fk)" json:"-"`
	WareHouseID int       `orm:"column(wh_id)" json:"wareHouseId"`
	Stock       int       `orm:"column(stock);null" json:"stock"`
	CreatedAt   time.Time `orm:"column(created_at);null" json:"createdAt"`
	UpdatedAt   time.Time `orm:"column(updated_at);null" json:"updatedAt"`
}

// TableName for users
func (pd *ProductDetails) TableName() string {
	return "products_details"
}
