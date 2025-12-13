package controllers

import (
	livenow "hltv/controllers/liveNow"
	"hltv/controllers/matches"

	"github.com/gin-gonic/gin"
)

func HLTVendpoints(router *gin.RouterGroup) {
	router.GET("/live-now", livenow.CallLiveNow)
	router.GET("/matches", matches.CallMatches)
}
