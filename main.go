package main

import (
	"fmt"
	"net/http"

	webart "webart/handler"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", webart.Asciiweb)
	fmt.Println("Server listening on port 7050:")
	fmt.Println("click http://localhost:7050")

	err := http.ListenAndServe(":7050", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
