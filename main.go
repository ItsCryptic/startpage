package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	fmt.Println("Server started on port 80")
	log.Fatal(http.ListenAndServe(":80", fs), nil)
}
