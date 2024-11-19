package main

import (
	"bpjs/confiq"
	"bpjs/routes"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := confiq.Connect()
	db.Close()
	r := routes.NewRouter()
	fmt.Println("Aplication Running in port :8080")
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
