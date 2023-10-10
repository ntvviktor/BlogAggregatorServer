package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func (apiConfig *apiConfig) HandleCreateFeedFollow(ctx *gin.Context) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	param := parameter{}
	err := ctx.BindJSON(&param)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Bad requested JSON")
		return
	}

	user := ctx.MustGet("user").(database.User)

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    param.FeedID,
		UserID:    user.ID,
	}

	createdFeed, err := apiConfig.DB.CreateFeedFollow(ctx.Request.Context(), newFeedFollow)
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, "Internal error")
		return
	}
	res := FeedFollow{
		ID:        createdFeed.ID,
		CreatedAt: createdFeed.CreatedAt,
		UpdatedAt: createdFeed.UpdatedAt,
		FeedID:    createdFeed.FeedID,
		UserID:    createdFeed.UserID,
	}
	respondWithJSON(ctx, http.StatusCreated, res)
}

func (apiConfig *apiConfig) HandleGetAllFeedFollowsByUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(database.User)
	feedFollows, err := apiConfig.DB.GetAllFeedFollowsByUser(ctx.Request.Context(), user.ID)
	if err != nil {
		respondWithError(ctx, http.StatusNotFound, "User or feeds not found")
		return
	}
	respondWithJSON(ctx, http.StatusOK, feedFollows)
}

func (apiConfig *apiConfig) HandleDeleteFeedFollow(ctx *gin.Context) {
	feedFollowID := ctx.Param("feedFollowID")
	str, _ := uuid.FromBytes([]byte(feedFollowID))

	user := ctx.MustGet("user").(database.User)

	err := apiConfig.DB.DeleteFeedFollow(ctx.Request.Context(), database.DeleteFeedFollowParams{
		ID:     str,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, "Cannot delete, internal error")
		return
	}
	res := map[string]string{
		feedFollowID: "deleted",
	}
	respondWithJSON(ctx, http.StatusOK, res)
}
