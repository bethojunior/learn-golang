package routes

import (
	"github.com/gin-gonic/gin"

	controllers "api/src/http/controllers"
)

var v1 *gin.RouterGroup

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	v1 = router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
	}

	return v1
}
