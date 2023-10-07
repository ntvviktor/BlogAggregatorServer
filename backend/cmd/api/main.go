package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ntvviktor/BlogApplication/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", dbUrl)
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	PORT := os.Getenv("PORT")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == fmt.Sprintf("http://localhost:%s", PORT)
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/v1")
	{
		v1.POST("/users", apiCfg.HandleCreateUsers)
		v1.GET("/users", apiCfg.HandleGetUserByID)
	}
	router.LoadHTMLGlob("templates/*")
	router.GET("/", renderHTML)
	router.Run(fmt.Sprintf(":%s", PORT))
}
