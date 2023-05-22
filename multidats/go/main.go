package main

import (
	"fmt"
	"time"

	"github.com/pai12345/Andromeda_Apps/multidats/go/customer"
	"github.com/pai12345/Andromeda_Apps/multidats/go/products"
	"github.com/pai12345/Andromeda_Apps/multidats/go/regions"
)

func main() {
	fmt.Println("=========Customer Data=========")
	user := customer.CustomerData{
		FirstName:   "Rahul",
		LastName:    "Pai",
		DateofBirth: "01-01-2000",
		Married:     true,
		PhoneNumber: "+1234",
		EmailAddr:   "abc@gmail.com",
		CreatedDate: time.Now(),
		Tags:        []string{"customer", "user"},
	}
	user.Generatedata()
	user.FindTag("customer")
	fmt.Println("=========Products Data=========")
	product := products.ProductsData{
		Name:        "asusIPRO5",
		Prodtype:    "Electronic",
		Description: "Asus IPRO 5 Series Laptop",
		Price:       "400",
		Currency:    "EUR",
		Tags:        []string{"laptop", "asus", "IPRO5"},
		ProductList: []string{"Electronic"},
	}
	product.Generatedata()
	product.AddProducts([]string{"Hardware", "Software", "Textile", "Grocery"})
	fmt.Println("=========Regions Data=========")
	allregions := regions.FetchAllregions()
	allregions.FetchRegion("Europe")
	err := allregions.CheckRegionsupport("Europ")
	if err != nil {
		fmt.Println(err)
	}
}
