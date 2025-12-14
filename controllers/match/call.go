package match

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMatchData(c *gin.Context) {
	MatchId := c.Query("matchid")
	MatchDescription := c.Query("matchDescription")
	if MatchId == "" || MatchDescription == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Required parameters missing",
		})
		return
	}
	maps, err := ExtractMatches()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, maps)
}
