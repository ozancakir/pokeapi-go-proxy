package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ozancakir/go-pokeapi-proxy/internals/api/dto"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db"
	"github.com/ozancakir/go-pokeapi-proxy/internals/db/entities"
)

type TranslateResponse struct {
	ResponseData struct {
		TranslatedText string  `json:"translatedText"`
		Match          float64 `json:"match"`
	} `json:"responseData"`
}

func Translate(c *gin.Context) {

	var req dto.TranslateRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	text := req.Text
	if text == "" {
		c.JSON(400, gin.H{
			"message": "text is required",
		})
		return
	}
	from := "en"
	if req.From != nil {
		from = *req.From
	}
	to := "tr"
	if req.To != nil {
		to = *req.To
	}

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.Trim(text, " ")

	host := os.Getenv("TRANSLATE_API_HOST")
	if host == "" {
		c.JSON(500, gin.H{
			"message": "TRANSLATE_API_HOST is not set",
		})
		return
	}
	path := os.Getenv("TRANSLATE_API_PATH")
	if path == "" {
		c.JSON(500, gin.H{
			"message": "TRANSLATE_API_PATH is not set",
		})
		return
	}
	user := os.Getenv("TRANSLATE_API_USER")
	if user == "" {
		c.JSON(500, gin.H{
			"message": "TRANSLATE_API_USER is not set",
		})
		return
	}
	pass := os.Getenv("TRANSLATE_API_KEY")
	if pass == "" {
		c.JSON(500, gin.H{
			"message": "TRANSLATE_API_KEY is not set",
		})
		return
	}

	db := db.GetDB()

	_u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	_q := _u.Query()
	_q.Add("q", text)
	_q.Add("de", user)
	_q.Add("key", pass)
	_q.Add("langpair", fmt.Sprintf("%s|%s", from, to))
	_u.RawQuery = _q.Encode()

	url := _u.String()

	var r entities.Translate
	tx := db.Where("url = ?", url).First(&r)
	if tx.Error == nil {
		c.JSON(200, gin.H{
			"result": r.Translation,
		})
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"message": "Translation API Error",
			"body":    string(body),
		})
		return
	}

	// read the response body and save to db

	var tr TranslateResponse
	err = json.Unmarshal(body, &tr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	r = entities.Translate{
		Url:         url,
		Translation: tr.ResponseData.TranslatedText,
	}
	go db.Create(&r)

	c.JSON(200, gin.H{
		"result": tr.ResponseData.TranslatedText,
	})

}
