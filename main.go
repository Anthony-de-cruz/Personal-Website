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

type Article struct {
	Header    string
	SubHeader string
	Sections  []Section
}

type Section struct {
	Header     string
	SubHeader  string
	Paragraphs []string
}

func loadArticle() Article {
	article := Article{
		Header:    "My Article",
		SubHeader: "A test article",
		Sections: []Section{
			{
				Header:    "Introduction",
				SubHeader: "how it all came to be . . .",
				Paragraphs: []string{
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce molestie vulputate arcu ac efficitur. Duis ac massa a est eleifend dignissim. Quisque lobortis vitae magna nec laoreet.",
					"Aenean at mollis massa, vel tincidunt odio. Nulla facilisi. Nam pellentesque risus arcu, vitae vulputate est commodo a. Nam eget feugiat tortor, quis posuere risus. Sed posuere magna vitae libero malesuada, id aliquam lacus facilisis.",
				},
			},
			{
				Header:    "Conclusion",
				SubHeader: ". . . in the end",
				Paragraphs: []string{
					"This is the end of the article.",
					"Suck on it.",
				},
			},
		},
	}

	return article
}

func renderTemplate(wr http.ResponseWriter, tmplName string) {
	err := templates.ExecuteTemplate(wr, tmplName+".html", loadArticle())
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
