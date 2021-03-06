package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars
func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range d {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := d[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}