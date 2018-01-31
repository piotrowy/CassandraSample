package main

import "net/http"

const (
	ip1 = ""
	ip2 = ""
	ip3 = ""
)

func main() {
	http.HandleFunc("/history", historyHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}
