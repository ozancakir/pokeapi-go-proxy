package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db/entities"
)

func ProxyPokeapi(c *gin.Context) {

	// get fullpath including query params
	fullpath := c.Request.URL.String()

	//check if we have the data in our database
	//if we have it, return it
	//if we don't have it, fetch it from pokeapi and save it to our database
	//then return it

	db := db.GetDB()

	var r entities.Response

	tx := db.Where("url = ?", fullpath).First(&r)

	if tx.Error == nil {
		// send r.Result as raw json
		c.Data(r.StatusCode, "application/json", []byte(r.Result))
		return

	}
	_prefix := os.Getenv("API_PREFIX")
	if _prefix == "" {
		_prefix = "/api"
	}
	pokeapi := os.Getenv("POKEAPI_URL")
	if pokeapi == "" {
		c.JSON(500, gin.H{
			"message": "POKEAPI_URL is not set",
		})
		return
	}

	proxyUri := fmt.Sprintf("%s%s", pokeapi, strings.Replace(fullpath, _prefix, "", 1))

	resp, err := http.Get(proxyUri)
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
		Url:        fullpath,
		Result:     string(body),
		StatusCode: resp.StatusCode,
	}

	go db.Create(&r)

	c.Data(200, "application/json", body)

}
