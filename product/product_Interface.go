package product

type ProductManager interface {
	Add(product *Product) error
	GetByID(id int) *Product
}
