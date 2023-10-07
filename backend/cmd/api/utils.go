package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getApiKey(ctx *gin.Context) string {
	apikeyPhrase := ctx.Request.Header.Get("Authorization")
	apikey := strings.TrimPrefix(apikeyPhrase, "ApiKey ")
	if apikey == apikeyPhrase {
		return ""
	}
	return apikey
}

func respondWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"error": message,
	})
}

func respondWithJSON(ctx *gin.Context, code int, res interface{}) {
	ctx.JSON(code, res)
}

func renderHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Mom",
	})
}
