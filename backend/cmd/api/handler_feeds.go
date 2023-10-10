package main

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

func (apiConfig *apiConfig) HandleCreateFeed(ctx *gin.Context) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	param := parameter{}
	err := ctx.BindJSON(&param)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Bad requested JSON")
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

	createdFeed, err := apiConfig.DB.CreateFeed(ctx.Request.Context(), newFeed)
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, "Internal error")
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
	respondWithJSON(ctx, http.StatusCreated, res)
}

func (apiConfig *apiConfig) HandleGetAllFeeds(ctx *gin.Context) {
	allFeeds, err := apiConfig.DB.GetAllFeeds(ctx.Request.Context())
	if err != nil {
		respondWithError(ctx, http.StatusNotFound, "Resources Not Found")
		return
	}
	respondWithJSON(ctx, http.StatusCreated, allFeeds)
}
