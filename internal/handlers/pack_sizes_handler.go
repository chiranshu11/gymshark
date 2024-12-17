package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gymshark/internal/storage"
)

type AddPackRequest struct {
	Size int `json:"size"`
}

func PackSizesHandler(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	switch r.Method {
	case http.MethodGet:
		getPackSizes(w, store)
	case http.MethodPost:
		addPackSize(w, r, store)
	case http.MethodDelete:
		deletePackSize(w, r, store)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPackSizes(w http.ResponseWriter, store storage.Storage) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.GetPackSizes())
}

func addPackSize(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	var req AddPackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Size <= 0 {
		http.Error(w, "Pack size must be a positive integer", http.StatusBadRequest)
		return
	}

	store.AddPackSize(req.Size)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.GetPackSizes())
}

func deletePackSize(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		http.Error(w, "Missing 'size' query parameter", http.StatusBadRequest)
		return
	}

	sizeVal, err := strconv.Atoi(sizeStr)
	if err != nil || sizeVal <= 0 {
		http.Error(w, "Size must be a positive integer", http.StatusBadRequest)
		return
	}

	err = store.RemovePackSize(sizeVal)
	if err != nil {
		http.Error(w, "Pack size not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.GetPackSizes())
}
