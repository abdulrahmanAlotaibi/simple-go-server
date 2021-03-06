package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func (app *application) home (w http.ResponseWriter , r *http.Request){
		if r.URL.Path != "/" {
			app.notFound(w)
			return
		}
		w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) showSnippet (w http.ResponseWriter , r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet (w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))

}