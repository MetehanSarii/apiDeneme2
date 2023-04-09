package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryID  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

var products = []Product{
	{
		ID:          1,
		ProductName: "Laptop",
		CategoryID:  1,
		UnitPrice:   8000.99,
	},
	{
		ID:          2,
		ProductName: "Tablet",
		CategoryID:  1,
		UnitPrice:   5000.99,
	},
	{
		ID:          3,
		ProductName: "Water",
		CategoryID:  2,
		UnitPrice:   0.99,
	},
}

var categories = []Category{
	{
		ID:           1,
		CategoryName: "Computer",
	},
	{
		ID:           2,
		CategoryName: "Beverages",
	},
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var produdcts []Product
	json.Unmarshal(bodyBytes, &produdcts)
	fmt.Println(produdcts)
}
func AddProduct() {
	product := Product{ID: 4, ProductName: "Telephone", CategoryID: 1, UnitPrice: 6000.0}
	jsonProduct, err := json.Marshal(product)
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	fmt.Println("Kaydedildi")

}
func main() {
	AddProduct()
	GetAllProducts()
}
