package repositories

import (
	"diskha-test/models"
	"github.com/astaxie/beego/orm"
	"strconv"
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

// GetTotalQtyPendingOrderByProductPubID func for find pending total product by product pubID
func (repo OrderRepository) GetTotalQtyPendingOrderByProductPubID(pID string) (count int, err error) {
	sql := `
		SELECT SUM(qty) AS qty
		FROM orders_details od 
		INNER JOIN orders o on o.pubid = od.order_pubid
		INNER JOIN products p ON p.pubid = od.product_pubid 
		WHERE o.status = 'PENDING' AND p.pubid = ?`

	var rs []orm.Params
	_, err = repo.db.Raw(sql, pID).Values(&rs)

	if rs[0]["qty"] != nil {
		count, _ = strconv.Atoi(rs[0]["qty"].(string))
	}

	return count, err
}
