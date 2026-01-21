package main

import (
	"log"
	"net/http"
)

func main() {

	// Servir les fichiers statiques (html, images, vidéos, css, js)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Serveur lancé sur http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
