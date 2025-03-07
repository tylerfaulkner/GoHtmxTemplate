package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"web-server/pkg/dotenv"
	"web-server/pkg/templates"

	"github.com/a-h/templ"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	dotenv.LoadEnv(".env")
	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.SetOutput(os.Stdout)
	mux := http.NewServeMux()

	cssFile, err := staticFiles.ReadFile("static/index.css")
	if err != nil {
		log.Println("Error reading CSS file:", err)
		return
	}

	mux.Handle("/", templ.Handler(templates.HomePage()))
	mux.Handle("/index.css", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(cssFile))
	}))

	// Start server with logging middleware to log unhandled routes
	log.Println("Starting server on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
