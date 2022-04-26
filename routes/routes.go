package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoknopak/financial-transactions-analysis/controllers"
)

func HandleRequest() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.GET("/", controllers.Index)

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", controllers.Upload)
	router.Run()
}
