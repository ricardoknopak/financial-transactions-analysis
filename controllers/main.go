package controllers

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/ricardoknopak/financial-transactions-analysis/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("transaction_file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	savedFile, err := os.Create(handler.Filename)
	defer savedFile.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(savedFile, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ReadCsv(handler.Filename)
}

func ReadCsv(filename string) {
	csvFile, err := os.Open(filename)
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
