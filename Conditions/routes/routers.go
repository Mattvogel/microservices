/*
 * API-Server
 *
 * This is the Swagger file for the Scale_Sanctuary API-Server
 *
 * API version: 1.0.0
 * Contact: matt@mattvogel.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routes

import (
	"net/http"

	"Conditions/controllers"

	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

var auth = new(controllers.AuthController)
var cond = new(controllers.ConditionController)

//TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

// CORS ...
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//RequestIDMiddleware ...
//Generate a unique ID and attach it to each request for future reference
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())
	router.Use(RequestIDMiddleware())

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, TokenAuthMiddleware(), route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, TokenAuthMiddleware(), route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, TokenAuthMiddleware(), route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, TokenAuthMiddleware(), route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"SendTemperature",
		http.MethodPost,
		"/v1/device/:deviceId/temp",
		cond.SendTemperature,
	},
	{
		"GetConditions",
		http.MethodGet,
		"/v1/device/:deviceID/conditions",
		cond.GetTemperature,
	},
}
