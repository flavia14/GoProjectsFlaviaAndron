package main

import (
	"fmt"
	"net/http"
	"app/pkg/handlers"
)

const portNumber = ":8084"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
