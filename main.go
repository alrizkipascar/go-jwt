package main

import (
	"log"
)

func main() {
	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}
	//
	// fmt.Printf("%v\n", store)

	if err := store.init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()
}
