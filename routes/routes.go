package routes

import (
	"net/http"

	"github.com/ricardoknopak/financial-transactions-analysis/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/upload", controllers.Upload)
}
