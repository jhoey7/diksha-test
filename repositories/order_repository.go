package repositories

import (
	"edot-test/models"
	"github.com/astaxie/beego/orm"
)

// OrderRepository struct
type OrderRepository struct {
	db orm.Ormer
}

// NewOrderRepository is func for initiate OrderRepository
func NewOrderRepository(o orm.Ormer) OrderRepository {
	return OrderRepository{db: o}
}

// Insert func for insert voucher
func (repo OrderRepository) Insert(o models.Orders) (models.Orders, error) {
	_, err := repo.db.Insert(&o)
	return o, err
}

func (repo OrderRepository) FindPendingOrderByPubId(pubID string) (models.Orders, error) {
	var order models.Orders
	err := repo.db.QueryTable("orders").
		Filter("pubid", pubID).
		Filter("status", "PENDING").
		One(&order)

	return order, err
}

// UpdateColumns function for update transaction data using certain columns
func (repo OrderRepository) UpdateColumns(p models.Orders, cols ...string) error {
	_, err := repo.db.Update(&p, cols...)
	return err
}
