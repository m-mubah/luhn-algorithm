package luhn_algorithm

import (
	"cryptosystem/luhn-algorithm/handlers"
	"github.com/gorilla/mux"
)

func AddLuhnAlgorithmRoutes(r *mux.Router) {
	//routes := r.PathPrefix("").Subrouter()

	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/validate", handlers.Validate).Methods("POST")
}
