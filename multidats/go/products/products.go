package products

import "fmt"

func (prods ProductsData) Generatedata() *ProductsData {
	fmt.Println(&prods)
	return &prods
}

func (prods ProductsData) AddProducts(prodlist []string) *[]string {
	prodlist = append(prods.ProductList, prodlist...)
	fmt.Println(&prodlist)
	return &prodlist
}
