package controller

import (
	"context"
	"net/http"
	"fmt"
	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
	"strconv"
)


func GetSmartLights(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
    return
	}

	if interval > 400{
		c.JSON(400, gin.H{"error": "Interval must be less than 400"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "SmartLights"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
					"data_humidity": value["data_humidity"],
					"data_temperature": value["data_temperature"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], 
	}       
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetSmartLightbyNodeName(c *gin.Context) {
	nodename := c.Param("nodename") // Parameter to query

	var objs = []gin.H{} // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "SmartLights"
		WHERE "nodeName" = '` + nodename + `'
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
					"data_humidity": value["data_humidity"],
					"data_temperature": value["data_temperature"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], 
	}       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}

	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceName not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

func GetSmartLightbyDevEUI(c *gin.Context) {
	devEUI := c.Param("devEUI") // Parameter to query
	var objs = []gin.H{}        // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "SmartLights"
		WHERE "devEUI" = '` + devEUI + `'
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { 
		value := iterator.Value() 
		obj := gin.H{
			"fields": gin.H{
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
					"data_humidity": value["data_humidity"],
					"data_temperature": value["data_temperature"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], 
	}       
		objs = append(objs, obj)  
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs) 
}
