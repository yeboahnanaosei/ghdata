package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbConnection *sql.DB

func loggingMiddleware(c *gin.Context) {
	msg := fmt.Sprintf("%v\t%v\t", c.Request.URL.RequestURI(), c.Request.UserAgent())
	log.Println(msg)
	c.Next()
}

func setup() (*gin.Engine, error) {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	// Setup logger
	logFile, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("could not open log file")
		return nil, errors.New("could not open log file")
	}
	log.SetOutput(logFile)

	// Setup database connection
	sqliteFilePath := os.Getenv("SQLITE")
	sqliteFilePath, _ = filepath.Abs(sqliteFilePath)
	dbConnection, err = sql.Open("sqlite3", sqliteFilePath)
	if err != nil {
		return nil, err
	}

	// Setup router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(loggingMiddleware)
	router.Use(location.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(routeNotFound)
	router.GET("/", handleIndex)

	// Setup region routes
	router.GET("regions", getAllRegions)
	router.GET("regions/:code", getOneRegion)

	// Setup district routes
	router.GET("districts", getAllDistricts)

	// Setup search routes
	router.GET("search/regions/:query", searchRegion)
	router.GET("search/districts/:query", searchDistrict)
	return router, nil
}

func main() {
	router, err := setup()
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9900"
	}

	fmt.Println("Running on port:", port)
	router.Run()
}
