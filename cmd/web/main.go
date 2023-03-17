package main

import (
	"fmt"
	"net/http"

	"github.com/paltmanndev/magic/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server auf Port%s gestartet\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
