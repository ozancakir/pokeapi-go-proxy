package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api/controller"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api/middleware"
)

func Setup(g *gin.Engine) {

	g.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "User-Agent", "x-api-key"},
			AllowCredentials: true,
			AllowWildcard:    true,
			AllowFiles:       true,
			AllowWebSockets:  true,

			// AllowAllOrigins:  true,
		}),
	)

	apiPrefix := os.Getenv("API_PREFIX")

	if apiPrefix == "" {
		apiPrefix = "/api"
	}
	api := g.Group(apiPrefix)
	// guard api with api key
	api.Use(middleware.ApiKeyAuth())

	//catch all routes
	api.POST("/translate", controller.Translate)
	api.GET("/*path", controller.ProxyPokeapi)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
