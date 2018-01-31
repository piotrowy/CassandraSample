package main

import (
	"time"
	"net/http"
)

type History struct {
	UserId       string
	ProductName  string
	ProductPrice int
	Date         time.Time
}

func selectHistory(userId string) []History {
	return []History{}
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userId := extURL(*r.URL).last()
		renderTemplate(w, "history", selectHistory(userId))
	}
}
