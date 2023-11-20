package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db"
	"github.com/ozancakir/go-pokeapi-proxy/utils"
)

func main() {

	err := godotenv.Load(utils.Path("./.env"))
	if err != nil {
		log.Panicln("Error loading .env file")

	}

	db.Setup()

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.New()
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())

	api.Setup(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalln("Failed to start server")
	}

	log.Println("Server started on port 4444")

}
