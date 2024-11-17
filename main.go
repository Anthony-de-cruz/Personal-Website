package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/anthony-de-cruz/Personal-Website/middleware"
)

const PORT uint16 = 8080

var templates = template.Must(
	template.ParseFiles(
		"views/index.html",
		"views/footer.html",
		"views/article.html",
	),
)

type Page struct {
	Footer *template.Template
}

func renderTemplate(wr http.ResponseWriter, tmplName string) {
	err := templates.ExecuteTemplate(wr, tmplName+".html", nil)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

func getArticleByName(wr http.ResponseWriter, req *http.Request) {
	// name := req.PathValue("name")
	renderTemplate(wr, "article")
}

func getHomePage(wr http.ResponseWriter, req *http.Request) {
	renderTemplate(wr, "index")
}

func handleNotFound(wr http.ResponseWriter, req *http.Request) {
	http.NotFound(wr, req)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", getHomePage)
	router.HandleFunc("GET /article", getArticleByName)

	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("GET /public/*", http.StripPrefix("/public/", fileServer))

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: middlewareStack((router)),
	}

	log.Printf("Starting server on port :%d\n", PORT)
	server.ListenAndServe()
}
