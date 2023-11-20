package controller

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db/entities"
)

func ProxyPokeapi(c *gin.Context) {
	path := c.Param("path")

	//check if we have the data in our database
	//if we have it, return it
	//if we don't have it, fetch it from pokeapi and save it to our database
	//then return it

	db := db.GetDB()

	var r entities.Response

	tx := db.Where("url = ?", path).First(&r)

	if tx.Error == nil {
		// send r.Result as raw json
		c.Data(r.StatusCode, "application/json", []byte(r.Result))
		return

	}

	pokeapi := os.Getenv("POKEAPI_URL")

	resp, err := http.Get(pokeapi + path)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// read the response body and save to db
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// replace all https://pokeapi.co/api/v2 with /api
	body = []byte(strings.ReplaceAll(string(body), pokeapi, "/api"))
	r = entities.Response{
		Url:        path,
		Result:     string(body),
		StatusCode: resp.StatusCode,
	}

	go db.Create(&r)

	c.Data(200, "application/json", body)

}
