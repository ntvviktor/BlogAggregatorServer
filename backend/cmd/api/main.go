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
	// TODO: automatically follow their own feed that is created by the users
	// feed, err := URLToFeed("https://feeds.feedburner.com/TechCrunch/")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}
	// Scape data feed from url
	// go startScraping(dbQueries, 3, time.Minute)

	PORT := os.Getenv("PORT")
	// Initialize router and configure CORS
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
		v1.GET("/feeds", apiCfg.HandleGetAllFeeds)
		v1.POST("/users", apiCfg.HandleCreateUsers)
		v1.Use(apiCfg.middlewareAuth())
		v1.GET("/users", apiCfg.HandleGetUserByID)
		v1.GET("/posts", apiCfg.HandleGetPostsByUser)
		v1.DELETE("/users/:feedFollowID", apiCfg.HandleDeleteFeedFollow)
		v1.POST("/feed_follows", apiCfg.HandleCreateFeedFollow)
		v1.POST("/feeds", apiCfg.HandleCreateFeed)
	}
	router.LoadHTMLGlob("templates/*")
	router.GET("/", renderHTML)
	router.Run(fmt.Sprintf(":%s", PORT))
}
