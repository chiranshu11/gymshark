package handlers

import (
	"net/http"

	"gymshark/internal/storage"
)

func RegisterRoutes(mux *http.ServeMux, store storage.Storage) {
	mux.HandleFunc("/", ServeUIHandler)

	mux.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		CalculateHandler(w, r, store)
	})

	mux.HandleFunc("/pack-sizes", func(w http.ResponseWriter, r *http.Request) {
		PackSizesHandler(w, r, store)
	})
}
