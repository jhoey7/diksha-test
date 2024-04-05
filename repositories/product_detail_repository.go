package repositories

import (
	"diskha-test/models"
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

func (repo ProductDetailRepository) FindByProductPubID(p string) ([]models.ProductDetails, error) {
	var pd []models.ProductDetails

	qs := repo.db.QueryTable("products_details")
	qs = qs.Filter("product_pubid", p)
	_, err := qs.All(&pd)

	return pd, err
}
