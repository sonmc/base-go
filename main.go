package main

import (
	"be/application/routes"

	"github.com/gin-gonic/gin"
)
var router = gin.Default()

func main() { 
	routes.Run( )
}