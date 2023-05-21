package products

import "fmt"

func (Products ProductsData) Generatedata() *ProductsData {
	fmt.Println(&Products)
	return &Products
}
