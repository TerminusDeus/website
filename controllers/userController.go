package controllers

import (
    "fmt"
    "html/template"
    "net/http"
	"website/models"
)

func IndexHandler(w http.ResponseWriter, r* http.Request) {	
	t, err := template.ParseFiles("templates/table.html")
	
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }   
	t.ExecuteTemplate(w, "table.html", models.InitData())	
}