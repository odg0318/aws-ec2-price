package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/odg0318/aws-ec2-price/pkg/price"
)

func getEc2PriceHandler(context *gin.Context) {
	region := context.Param("region")
	instanceType := context.Param("instance_type")

	pricing, err := price.NewPricing()
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	instance, err := pricing.GetInstance(region, instanceType)
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	context.JSON(200, instance)
}
