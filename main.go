package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(middleware.RedirectSlashes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		text := r.URL.Query().Get("text")
		if text != "" {
			fmt.Fprint(w, strings.ToUpper(text))
			return
		}

		fmt.Fprint(w, "Usage: /hello will return HELLO. /?text=hello will return HELLO")
	})

	r.Get("/{text}", func(w http.ResponseWriter, r *http.Request) {
		text, err := url.QueryUnescape(chi.URLParam(r, "text"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, strings.ToUpper(text))
	})

	http.ListenAndServe(os.Getenv("PORT"), r)
}
