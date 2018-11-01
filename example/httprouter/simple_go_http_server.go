package main

import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		log.Println("parse connect")
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8089", nil)
}