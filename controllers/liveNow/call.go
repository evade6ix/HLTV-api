package livenow

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallLiveNow(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Date is required",
		})
		return
	}
	fmt.Println("Iniciado")
	response, err := ExtractLiveNow(date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
