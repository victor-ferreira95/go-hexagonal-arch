package server

import (
	"go-hexagonal/adapters/web/handler"
	"go-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Webserver struct {
	service application.ProductServiceInterface
}

func NewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{service}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.NewProductHandlers(r, n, w.service)
	http.Handle("/", r)

	server := http.Server{
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
