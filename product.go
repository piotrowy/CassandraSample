package main

import (
	"net/http"
	"github.com/gocql/gocql"
)

type Product struct {
	Category     string
	Name  string
	Price int
}

func selectProducts(cat string) []Product {
	dataCh := make(chan []Product)

	WP.DoJob(func(session *gocql.Session) {
		iter := session.Query(ProductsSelect, cat).Consistency(gocql.Quorum).Iter()
		var (
			products []Product
			product_name string
			product_price int
		)
		for iter.Scan(&product_name, &product_price) {
			products = append(products, Product{
				Category: cat,
				Name: product_name,
				Price: product_price,
			})
		}
		dataCh <- products
	})

	return <-dataCh
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	category := extURL(*r.URL).last()
	renderTemplate(w, "products", selectProducts(category))
}
