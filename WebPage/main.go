package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantilla = template.Must(template.ParseGlob("Public/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	css := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", css))
	log.Println("Iniciando el servidor...")
	log.Println("Cargando las páginas...")
	log.Println("Servidor iniciado en http://localhost:8080/")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
