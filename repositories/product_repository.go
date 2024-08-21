package repositories

import (
	"edot-test/models"
	"github.com/astaxie/beego/orm"
)

// ProductRepository struct
type ProductRepository struct {
	db orm.Ormer
}

// NewProductRepository is func for initiate ProductRepository
func NewProductRepository(o orm.Ormer) ProductRepository {
	return ProductRepository{db: o}
}

// FindProductByQueryParam func
func (repo ProductRepository) FindProductByQueryParam(req models.ProductListRequest) (products []models.Products, err error) {
	o := repo.db.QueryTable("products")

	if req.Sort == "DESC" {
		o = o.OrderBy("-" + req.Order)
	} else {
		o = o.OrderBy(req.Order)
	}

	if req.Keyword != "" {
		o = o.Filter("product_name__icontains", req.Keyword)
	}

	_, err = o.Limit(req.Size).Offset(req.Page).All(&products)

	for i, p := range products {
		repo.db.LoadRelated(&p, "ProductDetail")
		products[i].ProductDetail = p.ProductDetail
	}

	return products, err
}

// CountProductByQueryParam function for find product by query param.
func (repo ProductRepository) CountProductByQueryParam(req models.ProductListRequest) (int, error) {
	o := repo.db.QueryTable("products")
	if req.Keyword != "" {
		o = o.Filter("product_name__icontains", req.Keyword)
	}

	count, err := o.Count()
	return int(count), err
}

// FindByPubID function for find product by pubID.
func (repo ProductRepository) FindByPubID(pubID string) (product models.Products, err error) {
	err = repo.db.QueryTable("products").Filter("pubid", pubID).One(&product)
	return product, err
}

// UpdateColumns function for update transaction data using certain columns
func (repo ProductRepository) UpdateColumns(p models.Products, cols ...string) error {
	_, err := repo.db.Update(&p, cols...)
	return err
}

// FindByPubIDs function for find product by pubIDs.
func (repo ProductRepository) FindByPubIDs(pubIDs []string) (products []models.Products, err error) {
	_, err = repo.db.QueryTable("products").
		Filter("pubid__in", pubIDs).
		All(&products)

	for i, p := range products {
		repo.db.LoadRelated(&p, "ProductDetail")
		products[i].ProductDetail = p.ProductDetail
	}
	return products, err
}
