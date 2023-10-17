package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ApiConfig *ApiConfig) MiddlewareAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := GetApiKey(ctx)
		if apikey == "" {
			RespondWithError(ctx, http.StatusUnauthorized, "Malformed Token")
			return
		}
		user, err := ApiConfig.DB.GetUserById(ctx.Request.Context(), apikey)
		if err != nil {
			RespondWithError(ctx, http.StatusNotFound, "No user found")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
