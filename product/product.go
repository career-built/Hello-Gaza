package product

import (
	"database/sql"
	"hello-gaza/models"
	"strconv"
)

// ProductManager ...
type ProductManager struct {
	dbConnector dbConnector
	sqlBuilder  sqlBuilder
}

func NewManager(connector dbConnector, builder sqlBuilder) *ProductManager {
	return &ProductManager{
		dbConnector: connector,
		sqlBuilder:  builder,
	}
}

type dbConnector interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type mongoConnector interface {
	Insert(query string, args ...any) (*sql.Rows, error)
}

type sqlBuilder interface {
	InsertBuild(tableName string, argsKeys []string, argsVals []string) string
}

func (p *ProductManager) Add(product *models.Product) error {

	// todo: insert into db product table
	priceStr := strconv.Itoa(product.PRICE)

	keys := []string{"name", "price"}
	vals := []string{
		product.NAME,
		priceStr,
	}
	_, err := p.dbConnector.Query(p.sqlBuilder.InsertBuild("product", keys, vals))
	return err
}
