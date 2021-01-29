package controllers

import (
	"net/http"
	"strconv"

	"Conditions/forms"
	"Conditions/models"

	"github.com/gin-gonic/gin"
)

// ConditionController ...
type ConditionController struct{}

var conditionModel = new(models.ConditionModel)

//SendTemperature - send temperature state from nodes to the Database
func (ctrl ConditionController) SendTemperature(c *gin.Context) {

	deviceID := c.Param("deviceId")
	var temperatureForm forms.Conditions
	c.Bind(&temperatureForm)
	device, err := conditionModel.SendTemperature(deviceID, temperatureForm)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success", "device": device})
	}
}

//GetTemperature ...
func (ctrl ConditionController) GetTemperature(c *gin.Context) {
	deviceID := c.Param("deviceId")
	timeform := c.PostForm("timeframe")
	timeframe, err := strconv.Atoi(timeform)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "please specify a proper timeframe in seconds"})
	} else {

		device, err := conditionModel.GetTemperature(deviceID, timeframe)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": device})
		}
	}
}
