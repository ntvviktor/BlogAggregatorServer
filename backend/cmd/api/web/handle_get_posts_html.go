package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ntvviktor/BlogApplication/cmd/api/handlers"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type Post struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	PubDate     string `json:"pubdate"`
}

func (apiConfig *Config) HandleGetPosts(ctx *gin.Context) {
	var limit int
	var err error
	limitQuery := ctx.DefaultQuery("limit", "unknown")
	if limitQuery == "unknown" {
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
		handlers.RespondWithError(ctx, 500, "Internal Error")
		return
	}
	var res []handlers.Post = handlers.ExportPosts(posts)
	ctx.HTML(http.StatusOK, "articles_lists.html", gin.H{
		"Posts": res,
	})
}
