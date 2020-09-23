package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeboahnanaosei/ghdata/db/sqlite"
)

func getAllDistricts(c *gin.Context) {
	service := sqlite.DistrictService{DB: dbConnection}
	districts, err := service.GetAllDistricts()
	if err != nil {
		c.JSON(http.StatusOK, payload{
			Success: false,
			Code:    http.StatusInternalServerError,
			Msg:     "An internal error occured",
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, payload{
		Success: true,
		Code:    http.StatusOK,
		Msg:     "Request successful",
		Data:    districts,
	})
}
