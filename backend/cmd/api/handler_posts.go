package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

func (apiConfig *apiConfig) HandleGetPostsByUser(ctx *gin.Context) {
	var limit int
	var err error
	limitQuery := ctx.DefaultQuery("limit", "unknow")
	if limitQuery == "unknow" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			limit = 10
			log.Println(err)
		}
	}

	user := ctx.MustGet("user").(database.User)
	posts, err := apiConfig.DB.GetPostsByUser(ctx.Request.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(ctx, 500, "Internal Error")
		return
	}
	respondWithJSON(ctx, http.StatusOK, posts)
}
