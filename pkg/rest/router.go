package rest

import (
	"github.com/gin-gonic/gin"
)

func errorHandler(context *gin.Context) {
	context.JSON(500, gin.H{"error": context.Errors.String()})
	context.Abort()
	context.Header("Content-Type", "application/json")
}

func notFoundHandler(context *gin.Context) {
	context.JSON(404, gin.H{"error": "no route"})
	context.Abort()
	context.Header("Content-Type", "application/json")
}

func GetRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(notFoundHandler)

	router.GET("/ec2/regions/:region", getEc2PricesHandler)
	router.GET("/ec2/regions/:region/instance_types/:instance_type", getEc2PriceHandler)

	return router
}
