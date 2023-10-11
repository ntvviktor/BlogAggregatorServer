package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (apiConfig *apiConfig) middlewareAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := getApiKey(ctx)
		if apikey == "" {
			respondWithError(ctx, http.StatusUnauthorized, "Malformed Token")
			return
		}
		user, err := apiConfig.DB.GetUserById(ctx.Request.Context(), apikey)
		if err != nil {
			respondWithError(ctx, http.StatusNotFound, "No user found")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
