package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account =>", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "anthony", "anthony", "anthony")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}
	//
	// fmt.Printf("%v\n", store)

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	// seed
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)

	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
