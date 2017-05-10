package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/tep-xi/tep-wiki/wiki", http.StatusSeeOther)
	})

	log.Println("listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
