package basket

import "modul/product"

type Basket struct {
	ProductList product.ProductList
	Total       int
}

func NewBasket() Basket {
	return Basket{}
}
