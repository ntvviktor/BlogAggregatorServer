package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntvviktor/BlogApplication/cmd/api/handlers"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type Config struct {
	*handlers.ApiConfig
}

func (apiConfig *Config) RenderHTML(ctx *gin.Context) {
	// allFeeds, err := ApiConfig.DB.GetAllFeeds(ctx.Request.Context())
	// if err != nil {
	// 	respondWithError(ctx, http.StatusNotFound, "Resources Not Found")
	// 	return
	// }
	// var res []Feed
	// for _, v := range allFeeds {
	// 	res = append(res, Feed{
	// 		ID:   v.ID,
	// 		Name: v.Name,
	// 	})
	// }

	var IsLogin bool = true
	apikey := handlers.GetApiKey(ctx)
	if apikey == "" {
		apikey = "6ab785413e828198848f189141ad26be20cb619d3812db118fb66b70948f5d2a"
		IsLogin = false
	}
	user, _ := apiConfig.DB.GetUserById(ctx.Request.Context(), apikey)
	posts, err := apiConfig.DB.GetPostsByUser(ctx.Request.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(15),
	})
	if err != nil {
		handlers.RespondWithError(ctx, 500, "Internal Error")
		return
	}

	var res []handlers.Post = handlers.ExportPosts(posts)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Posts":    res,
		"IsLogin":  IsLogin,
		"Username": user.Name,
	})
}
