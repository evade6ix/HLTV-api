package mapstats

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MapData struct {
	MatchId          string `json:"matchid"`
	MatchDescription string `json:"matchDescription"`
}

func GetMapStats(c *gin.Context) {
	var raw MapData
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := ExtractMatches(raw.MatchId, raw.MatchDescription)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
