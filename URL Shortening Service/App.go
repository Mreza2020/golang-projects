package main

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

var urlStore = make(map[string]string)

func generateCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, n)

	for i := range code {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		code[i] = letters[num.Int64()]
	}

	return string(code)
}

func shortenHandler(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	code := generateCode(6)
	urlStore[code] = req.URL

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + code,
	})
}

func redirectHandler(c *gin.Context) {
	code := c.Param("code")
	originalURL, exists := urlStore[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusFound, originalURL)
}

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenHandler)
	r.GET("/:code", redirectHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
