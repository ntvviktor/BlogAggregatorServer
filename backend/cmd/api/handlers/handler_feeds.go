package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (ApiConfig *ApiConfig) HandleCreateFeed(ctx *gin.Context) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	param := parameter{}
	err := ctx.BindJSON(&param)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, "Bad requested JSON")
		return
	}

	user := ctx.MustGet("user").(database.User)

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
		Url:       param.Url,
		UserID:    user.ID,
	}

	createdFeed, err := ApiConfig.DB.CreateFeed(ctx.Request.Context(), newFeed)
	if err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, "Internal error")
		return
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    createdFeed.ID,
		UserID:    user.ID,
	}

	_, err = ApiConfig.DB.CreateFeedFollow(ctx.Request.Context(), newFeedFollow)
	if err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, "Internal error")
		return
	}

	res := Feed{
		ID:        createdFeed.ID,
		CreatedAt: createdFeed.CreatedAt,
		UpdatedAt: createdFeed.UpdatedAt,
		Name:      createdFeed.Name,
		Url:       createdFeed.Url,
		UserID:    createdFeed.UserID,
	}
	RespondWithJSON(ctx, http.StatusCreated, res)
}

func (ApiConfig *ApiConfig) HandleGetAllFeeds(ctx *gin.Context) {
	allFeeds, err := ApiConfig.DB.GetAllFeeds(ctx.Request.Context())
	if err != nil {
		RespondWithError(ctx, http.StatusNotFound, "Resources Not Found")
		return
	}
	RespondWithJSON(ctx, http.StatusCreated, allFeeds)
}
