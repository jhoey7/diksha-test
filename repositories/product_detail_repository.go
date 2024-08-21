package repositories

import (
	"edot-test/models"
	"github.com/astaxie/beego/orm"
)

// ProductDetailRepository struct
type ProductDetailRepository struct {
	db orm.Ormer
}

// NewProductDetailRepository is func for initiate ProductDetailRepository
func NewProductDetailRepository(o orm.Ormer) ProductDetailRepository {
	return ProductDetailRepository{db: o}
}

func (repo ProductDetailRepository) FindByPubIds(ids []int) ([]models.ProductDetails, error) {
	var pd []models.ProductDetails

	_, err := repo.db.QueryTable("products_details").
		Filter("id__in", ids).
		All(&pd)

	return pd, err
}

// UpdateColumns function for update transaction data using certain columns
func (repo ProductDetailRepository) UpdateColumns(p models.ProductDetails, cols ...string) error {
	_, err := repo.db.Update(&p, cols...)
	return err
}
