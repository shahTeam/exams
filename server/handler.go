package server

import "github.com/gin-gonic/gin"

type HandlerRequests interface {
	TakeWord(c *gin.Context)
	GiveAllWord(c *gin.Context)
	GiveLimitedWord(c *gin.Context)
}

