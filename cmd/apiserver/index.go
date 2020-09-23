package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func handleIndex(c *gin.Context) {
	appURL := os.Getenv("BASE_URL")
	if appURL == "" {
		u := location.Get(c)
		appURL = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	}
	c.JSON(http.StatusOK, gin.H{
		"all_regions":                fmt.Sprintf("%s/regions", appURL),
		"all_regions_with_districts": fmt.Sprintf("%s/regions?embed=districts", appURL),
		"all_districts":              fmt.Sprintf("%s/districts", appURL),
		"one_region":                 fmt.Sprintf("%s/regions/region_code", appURL),
		"one_region_with_districts":  fmt.Sprintf("%s/regions/region_code?embed=districts", appURL),
		"search_regions":             fmt.Sprintf("%s/search/regions/:keyword", appURL),
	},
	)
}
