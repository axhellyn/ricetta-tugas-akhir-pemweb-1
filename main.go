package main

import (
	"log"
	"net/http"
	"ricetta/config"
	"ricetta/controllers/detailcontroller"
	"ricetta/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	//Homepage
	// http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/detail-recipe", detailcontroller.DetailRecipe)

	//detail recipe

	// http.HandleFunc("/detail", homecontroller.detail)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}