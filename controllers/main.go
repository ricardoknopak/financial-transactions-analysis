package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoknopak/financial-transactions-analysis/models"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Upload(c *gin.Context) {
	file, error := c.FormFile("transaction_file")
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	c.SaveUploadedFile(file, "uploads/"+file.Filename)
	ReadCsv(file.Filename)
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"Upload_result": "uploaded file: " + file.Filename,
	})
}

func ReadCsv(filename string) {
	csvFile, err := os.Open("uploads/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		valorTransacao, _ := strconv.ParseFloat(line[6], 64)
		transactions := models.Transactions{
			BancoOrigem:    line[0],
			AgenciaOrigem:  line[1],
			ContaOrigem:    line[2],
			BancoDestino:   line[3],
			AgenciaDestino: line[4],
			ContaDestino:   line[5],
			ValorTransacao: valorTransacao,
		}
		fmt.Println(transactions.BancoOrigem + " " + transactions.AgenciaOrigem + " " + transactions.ContaOrigem + " " + transactions.BancoDestino + " " + transactions.AgenciaDestino + " " + transactions.ContaDestino + " ")
	}
}
