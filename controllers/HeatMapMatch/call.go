package heatmapmatch

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MapData struct {
	MatchId          string `json:"matchid"`
	MatchDescription string `json:"matchDescription"`
}

func CallHeatMapMatch(c *gin.Context) {
	var raw MapData
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := GetHeatMap(raw.MatchId, raw.MatchDescription)
	c.Data(
		200,
		"image/png",
		data,
	)
}
