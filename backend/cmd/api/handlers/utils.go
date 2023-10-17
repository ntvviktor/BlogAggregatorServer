package handlers

import (
	"fmt"
	"html"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type Post struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	PubDate     string `json:"pubdate"`
}

func GetApiKey(ctx *gin.Context) string {
	apikeyPhrase := ctx.Request.Header.Get("Authorization")
	apikey := strings.TrimPrefix(apikeyPhrase, "ApiKey ")
	fmt.Println(apikey)
	if apikey == apikeyPhrase {
		return ""
	}
	return apikey
}

func RespondWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"error": message,
	})
}

func RespondWithJSON(ctx *gin.Context, code int, res interface{}) {
	ctx.JSON(code, res)
}

func ExportPosts(posts []database.GetPostsByUserRow) []Post {
	var res []Post
	for _, post := range posts {
		res = append(res, Post{
			Title:       post.Title,
			Url:         post.Url,
			Description: html.UnescapeString(post.Description.String),
			PubDate:     post.PublishedAt.Time.Format("02, Jan 2006"),
		})
	}
	return res
}
