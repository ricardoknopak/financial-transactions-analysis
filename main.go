package main

import (
	"net/http"

	"github.com/ricardoknopak/financial-transactions-analysis/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
