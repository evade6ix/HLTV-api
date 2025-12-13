package livenow

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CallLiveNow(c *gin.Context) {
	now := fmt.Sprintf("%v-%v-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	response, err := ExtractLiveNow(now)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
