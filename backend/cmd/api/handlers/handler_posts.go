package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

func (ApiConfig *ApiConfig) HandleGetPostsByUser(ctx *gin.Context) {
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
	posts, err := ApiConfig.DB.GetPostsByUser(ctx.Request.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		RespondWithError(ctx, 500, "Internal Error")
		return
	}
	RespondWithJSON(ctx, http.StatusOK, posts)
}
