package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/odg0318/aws-ec2-price/pkg/price"
)

func getEc2PricesHandler(context *gin.Context) {
	region := context.Param("region")

	pricing, err := price.NewPricing()
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	instances, err := pricing.GetInstances(region)
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	context.JSON(200, instances)
}
