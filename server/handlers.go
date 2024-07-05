package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Hello, there %s", "visitor")
	}
}

func getContries(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", contries)
}

func addContry(w http.ResponseWriter, r *http.Request) {

	country := &Contry{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	contries = append(contries, country)
	fmt.Fprintf(w, "country was added")
}
