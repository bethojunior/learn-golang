package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// _ entities "api/src/http/entities"
	types "api/src/types/public"
)

type tweetController struct {
	tweets []types.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (c *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.tweets)
}