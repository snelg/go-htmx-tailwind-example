package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//go:embed css/output.css
	css embed.FS

	//parsed templates
	templates *template.Template
)

func main() {

	//parse templates
	var err error
	templates, err = TemplateParseFSRecursive(templateFS, ".html")
	if err != nil {
		panic(err)
	}

	//add routes
	router := http.NewServeMux()
	router.Handle("GET /css/output.css", http.FileServer(http.FS(css)))

	router.Handle("GET /{$}", HtmlHandler(index))

	router.Handle("GET /companies", HtmlHandler(companies))

	router.Handle("GET /company/{id}", HtmlHandler(companyGet))
	router.Handle("PUT /company/{id}", HtmlHandler(companyPut))
	router.Handle("POST /company", HtmlHandler(companyPost))
	router.Handle("DELETE /company/{id}", HtmlHandler(companyDelete))
	router.Handle("GET /company/{id}/edit", HtmlHandler(companyEdit))
	router.Handle("GET /company/add", HtmlHandler(companyAdd))

	//logging/tracing
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := tracing(nextRequestID)(logging(logger)(router))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Println("listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, middleware); err != nil {
		logger.Println("http.ListenAndServe():", err)
		os.Exit(1)
	}
}

func TemplateParseFSRecursive(
	templates fs.FS,
	ext string) (*template.Template, error) {

	root := template.New("")
	err := fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ext) {
			if err != nil {
				return err
			}
			b, err := fs.ReadFile(templates, path)
			if err != nil {
				return err
			}
			parts := strings.Split(path, "/")
			name := strings.Join(parts[1:], "/")
			t := root.New(name)
			_, err = t.Parse(string(b))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return root, err
}
