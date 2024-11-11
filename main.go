package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize router
	router := gin.Default()
	// add all routes for service
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Path, route.Handler)
		case "POST":
			router.POST(route.Path, route.Handler)
		case "PUT":
			router.PUT(route.Path, route.Handler)
		case "PATCH":
			router.PATCH(route.Path, route.Handler)
		case "DELETE":
			router.DELETE(route.Path, route.Handler)
		case "OPTIONS":
			router.OPTIONS(route.Path, route.Handler)
		}
	}
	// run server
	router.Run("localhost:8080")
}
