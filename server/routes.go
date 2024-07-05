package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/contries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			getContries(w, r)

		case http.MethodPost:
			addContry(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not Allowed")
			return

		}

	})
}
