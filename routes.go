package main

import (
	"fetch-assessment/service"
	"github.com/gin-gonic/gin"
)

type route struct {
	Method string
	Path string
	Handler gin.HandlerFunc
}

var routes = []route{
	addRoute("GET", "/receipts/:id/points", service.GetReceipt),
	addRoute("POST", "/receipts/process", service.PostReceipt),
}

func addRoute(method string, path string, handler gin.HandlerFunc) route {
	return route{method, path, handler}
}
