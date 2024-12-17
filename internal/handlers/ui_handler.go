package handlers

import (
	"net/http"
	"path/filepath"
)

// ServeUIHandler serves the index.html from the static directory.
func ServeUIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// If you want to handle 404 for other paths that are not `/`, you can do so here.
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, filepath.Join("static", "index.html"))
}
