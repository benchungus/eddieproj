package main

import (
	"github.com/benchungus/api/handler"
	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	//load all env vars
	utils.GetEnvVar()

	//create router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	//set up endpoints
	router.GET("/stocks", handler.GetStocks)
	router.GET("/stocks/:name", handler.GetStockByName)
	router.GET("/bonds", handler.GetBond)
	router.GET("/stockCatalog/:name", handler.GrabStockFromCatalog)
	router.GET("/stocks/all", handler.AggBeta)
	router.POST("/bonds/:qty", handler.ChangeBond)
	router.POST("/stocks/:name/:qty", handler.NewStock)
	router.POST("/stocksMod/:name", handler.ChangeQty)
	router.DELETE("/stocks/:name", handler.DeleteStock)

	//establish connection
	router.Run("localhost:8080")
}
