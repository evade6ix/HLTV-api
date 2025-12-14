package results

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallResults(c *gin.Context) {
	response, err := ExtractResults()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}
