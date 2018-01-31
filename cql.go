package main

import "github.com/gocql/gocql"

const (
	ProductsInsert       = "INSERT INTO products (category, product_name, product_price) VALUES (?, ?, ?);"
	BuyCountersUpdate    = "UPDATE buy_counters set occurrences = occurrences + 1 where product_first = ? and product_second = ?;"
	HistoryInsert        = "INSERT INTO history (user_id, product_name, product_price, date) VALUES (?, ?, ?, ?);"
	RecommendationInsert = "INSERT INTO recommendation (product_name, other_product_name, other_product_price) VALUES (?, ?, ?);"
	OccurencesSelect     = "SELECT occurrences FROM buy_counters WHERE product_first = ? and product_second = ?;"
	RecommendationSelect = "SELECT * FROM recommendation WHERE product_name = ?;"
	ReccomendationDelete = "DELETE FROM recommendation WHERE product_name = ? and other_product_name = ?;"
	HistorySelect        = "SELECT * FROM history WHERE user_id = ?;"
	ProductsSelect       = "SELECT * FROM products WHERE category = ?;"
)

