package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gymshark/internal/storage"
	"gymshark/internal/usecase"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	itemsStr := r.URL.Query().Get("items")
	if itemsStr == "" {
		http.Error(w, "Missing 'items' query parameter", http.StatusBadRequest)
		return
	}
	items, err := strconv.Atoi(itemsStr)
	if err != nil || items < 1 {
		http.Error(w, "'items' must be a positive integer", http.StatusBadRequest)
		return
	}

	packs := store.GetPackSizes()
	result, err := usecase.CalculatePacks(items, packs)
	if err != nil {
		http.Error(w, "No solution found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
