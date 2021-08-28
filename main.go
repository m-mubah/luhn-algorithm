package main

import (
	luhnAlgorithm "cryptosystem/luhn-algorithm"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	csrfMiddleware := csrf.Protect([]byte("n&@ix77r#^&^cgeb13w@!+pht^6qu-=("),
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"*"}))

	luhnAlgorithm.AddLuhnAlgorithmRoutes(r)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	r.Use(csrfMiddleware)

	server := &http.Server{
		Addr:         ":8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	log.Fatal(server.ListenAndServe())
}
