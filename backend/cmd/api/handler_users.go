package main

import (
	"net/http"
	"time"

	"github.com/ntvviktor/BlogApplication/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func (apiConfig *apiConfig) HandleCreateUsers(ctx *gin.Context) {
	type parameter struct {
		Name string `json:"name"`
	}
	param := parameter{}
	err := ctx.BindJSON(&param)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Bad requested JSON")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	}

	createdUser, err := apiConfig.DB.CreateUser(ctx.Request.Context(), newUser)
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, "Cannot create users")
	}
	res := User{
		ID:        createdUser.ID,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		Name:      createdUser.Name,
		ApiKey:    createdUser.ApiKey,
	}
	respondWithJSON(ctx, http.StatusCreated, res)
}

func (apiConfig *apiConfig) HandleGetUserByID(ctx *gin.Context) {
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
	res := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
	respondWithJSON(ctx, http.StatusOK, res)
}
