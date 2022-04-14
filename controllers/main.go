package controllers

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
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
	fmt.Println("Worked!")
	readFile(handler.Filename)
}

func readFile(filename string) []string {
	var reading []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		reading = append(reading, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	fmt.Println(reading)
	return reading
}
