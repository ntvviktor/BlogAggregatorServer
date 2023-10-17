package handlers

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

func (ApiConfig *ApiConfig) HandleCreateUsers(ctx *gin.Context) {
	type parameter struct {
		Name string `json:"name"`
	}
	param := parameter{}
	err := ctx.BindJSON(&param)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, "Bad requested JSON")
		return
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	}

	createdUser, err := ApiConfig.DB.CreateUser(ctx.Request.Context(), newUser)
	if err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, "Cannot create users")
		return
	}
	res := User{
		ID:        createdUser.ID,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		Name:      createdUser.Name,
		ApiKey:    createdUser.ApiKey,
	}
	RespondWithJSON(ctx, http.StatusCreated, res)
}

func (ApiConfig *ApiConfig) HandleGetUserByID(ctx *gin.Context) {
	user := ctx.MustGet("user").(database.User)
	res := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
	RespondWithJSON(ctx, http.StatusOK, res)
}
