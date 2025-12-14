package match

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Match struct {
	MatchId          string `json:"matchid"`
	MatchDescription string `json:"matchDescription"`
}

func GetMatchData(c *gin.Context) {
	var raw Match
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := ExtractData(raw.MatchId, raw.MatchDescription)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
