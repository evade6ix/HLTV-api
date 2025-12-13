package matches

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallMatches(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "'date' is required",
		})
		return
	}

	response, err := ExtractMatches(date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
