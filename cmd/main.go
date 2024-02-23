package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	router "github.com/alrizkipascar/go-jwt/internal/api/routers"
	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/utils"
	"github.com/rs/cors"
)

func main() {
	// 8080
	seed := flag.Bool("seed", false, "seed the db")
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	//Port definition
	port := 8080
	//Start server
	log.Println("JSON API server running on port: ", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}
}
