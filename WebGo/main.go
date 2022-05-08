package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantilla = template.Must(template.ParseGlob("Public/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	_ = plantilla.ExecuteTemplate(w, "index.html", nil)
}
func info(w http.ResponseWriter, r *http.Request) {
	_ = plantilla.ExecuteTemplate(w, "info.html", nil)
}
func main() {
	css := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", css))
	log.Println("Iniciando el servidor...")
	log.Println("Cargando las páginas...")
	log.Println("Servidor iniciado en http://localhost:8080/")
	http.HandleFunc("/", index)
	http.HandleFunc("/informacion", info)
	_ = http.ListenAndServe(":8080", nil)
}
