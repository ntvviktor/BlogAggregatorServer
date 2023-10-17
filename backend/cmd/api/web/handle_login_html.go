package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (apiConfig *Config) HandleLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"successful": true,
	})
}
