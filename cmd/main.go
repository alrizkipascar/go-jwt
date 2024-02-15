package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	router "github.com/alrizkipascar/go-jwt/internal/api/routers"
	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/utils"
)

func main() {
	// 8080
	seed := flag.Bool("seed", true, "seed the db")
	flag.Parse()

	//Initialize database
	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	//seeding is allowed?
	if *seed {
		fmt.Println("seeding the database")
		utils.SeedAccounts()
	}
	//Routes definition
	router := router.InitRouter()
	//Port definition
	port := 8080
	//Start server
	log.Println("JSON API server running on port: ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}
}
