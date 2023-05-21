package main

import (
	"time"

	"github.com/pai12345/Andromeda_Apps/multidats/go/customer"
	"github.com/pai12345/Andromeda_Apps/multidats/go/products"
)

func main() {
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
	product := products.ProductsData{
		Name:        "asusIPRO5",
		Prodtype:    "Electronic",
		Description: "Asus IPRO 5 Series Laptop",
		Price:       "400",
		Currency:    "EUR",
		Tags:        []string{"laptop", "asus", "IPRO5"},
	}
	product.Generatedata()
}
