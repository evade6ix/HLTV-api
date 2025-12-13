package controllers

import (
	livenow "hltv/controllers/liveNow"

	"github.com/gin-gonic/gin"
)

func HLTVendpoints(router *gin.RouterGroup) {
	router.GET("/live-now", livenow.CallLiveNow)
}
