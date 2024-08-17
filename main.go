package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/asciiweb"
)

func main() {
	fmt.Println("server is running")
	http.HandleFunc("/", asciiweb.ProcessMainPage)
	fs := http.FileServer(http.Dir("Style"))
	http.Handle("/Style/", http.StripPrefix("/Style/", fs))

	http.HandleFunc("/ascii-art", asciiweb.HandleInput)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
