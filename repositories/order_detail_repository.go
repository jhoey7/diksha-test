package repositories

import (
	"edot-test/models"
	"github.com/astaxie/beego/orm"
)

// OrderDetailRepository struct
type OrderDetailRepository struct {
	db orm.Ormer
}

// NewOrderDetailRepository is func for initiate OrderDetailRepository
func NewOrderDetailRepository(o orm.Ormer) OrderDetailRepository {
	return OrderDetailRepository{db: o}
}

func (repo OrderDetailRepository) InsertBulk(ods []models.OrdersDetail) ([]models.OrdersDetail, error) {
	_, err := repo.db.InsertMulti(len(ods), ods)
	return ods, err
}
