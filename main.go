package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/timeseries/v0.2/smartcampusmaua")
	{
		api.GET("SmartLights", controller.GetSmartLights)
		api.GET("SmartLights/deviceName/:nodename", controller.GetSmartLightbyNodeName)
		api.GET("SmartLights/deviceId/:devEUI", controller.GetSmartLightbyDevEUI)

		api.GET("WaterTankLevel", controller.GetWaterTankLevel)
		api.GET("WaterTankLevel/deviceName/:nodename", controller.GetWaterTankLevelbyNodeName)
		api.GET("WaterTankLevel/deviceId/:devEUI", controller.GetWaterTankLevelbyDevEUI)
	}

	r.Run(":8888")
}
