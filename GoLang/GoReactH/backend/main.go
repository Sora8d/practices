package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func basic_route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, err := json.Marshal(albums[0:])
		if err != nil {
			log.Fatal(err)
		}
		//These 2 both work
		w.Write(res)
		//fmt.Fprintln(w, string(res))
	}
}

func main() {
	fmt.Println("Backend App v0.01")
	http.HandleFunc("/api/", basic_route)
	http.ListenAndServe(":8081", nil)
}
