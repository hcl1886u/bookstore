package main

import (
	"net/http"
	"text/template"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	t.Execute(w, "")
}
func main() {
	http.HandleFunc("/main", Indexhandler)
	http.ListenAndServe(":8081", nil)
}
