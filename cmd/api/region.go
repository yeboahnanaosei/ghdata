package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yeboahnanaosei/ghdata/db/sqlite"
)

func getAllRegions(c *gin.Context) {
	regionService := sqlite.RegionService{DB: dbConnection}
	regions, err := regionService.GetAllRegions()
	if err != nil {
		c.JSON(http.StatusOK, payload{
			Success: false,
			Code:    http.StatusInternalServerError,
			Msg:     "An internal error occured",
		})
		log.Println(err)
		return
	}

	embed := strings.ToLower(c.Query("embed"))
	if embed == "districts" {
		districtService := sqlite.DistrictService{DB: dbConnection}
		for _, r := range regions {
			districts, err := districtService.GetDistrictsByRegion(strings.ToUpper(r.Code))
			if err != nil {
				c.JSON(http.StatusOK, payload{
					Success: false,
					Code:    http.StatusInternalServerError,
					Msg:     "An internal error occured",
				})
				log.Println(err)
				return
			}
			r.Districts = districts
		}
	}

	c.JSON(http.StatusOK, payload{
		Success: true,
		Code:    http.StatusOK,
		Msg:     "Request successful",
		Data:    regions,
	})
}

func getOneRegion(c *gin.Context) {
	regionCode := c.Param("code")
	if regionCode == "" {
		c.JSON(http.StatusOK, payload{
			Success: false,
			Code:    http.StatusBadRequest,
			Msg:     "No region code supplied",
		})
		return
	}
	service := sqlite.RegionService{DB: dbConnection}
	region, err := service.GetOneRegion(strings.ToUpper(regionCode))
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusOK, payload{
			Success: false,
			Code:    http.StatusNotFound,
			Msg:     fmt.Sprintf("No region found with code %s", regionCode),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, payload{
			Success: false,
			Code:    http.StatusInternalServerError,
			Msg:     "An internal error occured",
		})
		log.Println(err)
		return
	}

	districtService := sqlite.DistrictService{DB: dbConnection}

	// Handle if the request includes an embed parameter
	embed := strings.ToLower(c.Query("embed"))
	if embed == "districts" {
		districts, err := districtService.GetDistrictsByRegion(strings.ToUpper(regionCode))
		if err != nil {
			c.JSON(http.StatusOK, payload{
				Success: false,
				Code:    http.StatusInternalServerError,
				Msg:     "An internal error occured",
			})
			log.Println(err)
			return
		}

		for _, district := range districts {
			region.Districts = append(region.Districts, district)
		}
	}

	c.JSON(http.StatusOK, payload{
		Success: true,
		Code:    http.StatusOK,
		Msg:     "Request successful",
		Data:    region,
	})
}
