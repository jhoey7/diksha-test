package models

import "time"

// Warehouse struct for warehouse
type Warehouse struct {
	ID          int       `orm:"auto;pk;column(id)" json:"id"`
	Code        string    `orm:"column(wh_code);null" json:"code"`
	Name        string    `orm:"column(wh_name);null" json:"name"`
	Description string    `orm:"column(wh_description);null" json:"description"`
	CreatedAt   time.Time `orm:"column(created_at);null" json:"createdAt"`
	UpdatedAt   time.Time `orm:"column(updated_at);null" json:"updatedAt"`
}

// TableName for users
func (w *Warehouse) TableName() string {
	return "warehouse"
}
