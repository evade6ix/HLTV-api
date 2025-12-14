package controllers

import (
	livenow "hltv/controllers/liveNow"
	mapstats "hltv/controllers/mapStats"
	"hltv/controllers/match"
	"hltv/controllers/matches"
	"hltv/controllers/results"

	"github.com/gin-gonic/gin"
)

func HLTVendpoints(router *gin.RouterGroup) {
	router.GET("/live-now", livenow.CallLiveNow)
	router.GET("/matches", matches.CallMatches)
	router.GET("/last-results", results.CallResults)
	router.POST("/match", match.GetMatchData)
	router.POST("/match-stats", mapstats.GetMapStats)
}
