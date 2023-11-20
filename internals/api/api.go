package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api/controller"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api/middleware"
)

func Setup(g *gin.Engine) {

	apiPrefix := os.Getenv("API_PREFIX")

	if apiPrefix == "" {
		apiPrefix = "/api"
	}
	api := g.Group(apiPrefix)
	// guard api with api key
	api.Use(middleware.ApiKeyAuth())

	//catch all routes
	api.GET("/*path", controller.ProxyPokeapi)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
