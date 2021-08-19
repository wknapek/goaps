package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port to start server")
	filename := flag.String("file", "gopher.json", "the JSON file with story")
	flag.Parse()
	fmt.Printf("parse file: %s\n", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	story, err := JsonStory(file)
	if err != nil {
		fmt.Println(err)
	}

	h := NewHandler(story, nil)
	fmt.Printf("starting on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
