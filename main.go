// website project main.go
package main

import (
    "net/http"
	"website/controllers"
)

func main() {
    http.HandleFunc("/", controllers.IndexHandler)	
    http.ListenAndServe(":8080", nil)
}

