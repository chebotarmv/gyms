package handlers

import "net/http"

func ServeUI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/index.html")
}
