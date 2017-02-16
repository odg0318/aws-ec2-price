package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/odg0318/aws-ec2-price/pkg/price"
)

type PriceResponse struct {
	Price        string `json:"price"`
	Region       string `json:"region"`
	InstanceType string `json:"instance_type"`
}

func getEc2PriceHandler(context *gin.Context) {
	region := context.Param("region")
	instanceType := context.Param("instance_type")

	pricing, err := price.NewPricing()
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	price, err := pricing.GetPrice(region, instanceType)
	if err != nil {
		context.Error(err)
		errorHandler(context)
		return
	}

	response := PriceResponse{price, region, instanceType}

	context.JSON(200, response)
}
