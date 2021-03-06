/*
 * API-Server
 *
 * This is the Swagger file for the Scale_Sanctuary API-Server
 *
 * API version: 1.0.0
 * Contact: matt@mattvogel.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controllers

import (
	"net/http"

	"Devices/forms"

	"Devices/models"

	"github.com/gin-gonic/gin"
)

// DeviceController ...
type DeviceController struct{}

var deviceModel = new(models.DeviceModel)

//getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

// AddDevice - Add a new device to the service
func (ctrl DeviceController) AddDevice(c *gin.Context) {

	var deviceForm forms.Device

	if c.ShouldBindJSON(&deviceForm) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid Form"})
		return
	}
	userID := getUserID(c)
	device, err := deviceModel.CreateNewDevice(userID, deviceForm)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	if device.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully created", "device": device})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not create this Device", "error": err.Error()})
	}
}

// UpdateDevice - Updates the owner of a device
func (ctrl DeviceController) UpdateDevice(c *gin.Context) {

	var deviceForm forms.Device
	device, err := deviceModel.UpdateDevice(deviceForm)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"device": device})
}

// DeleteDevice - Deletes a device
func DeleteDevice(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "This is the Delete Device placeholder, please fix"})
}

// GetDeviceByID - Find device by ID
func (ctrl DeviceController) GetDeviceByID(c *gin.Context) {

	deviceID := c.Param("deviceId")
	device, err := deviceModel.GetDeviceByID(deviceID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	if device.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"device": device})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not find this device", "error": err.Error()})
	}
}

// GetDeviceByOwner - Find devices by OwnerID
func (ctrl DeviceController) GetDeviceByOwner(c *gin.Context) {
	userID := getUserID(c)
	devices, err := deviceModel.GetDevicesByOwner(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"device": devices})
	}

}
