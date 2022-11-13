package handler

import (
	"net/http"

	"github.com/benchungus/api/db"
	"github.com/benchungus/api/models"

	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

func GrabStockFromCatalog(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("grabbing stock " + name)
	stock := db.GrabCatalogStock(name)
	utils.Logger.Info("got stock " + name)
	c.IndentedJSON(http.StatusOK, stock)
}

func GrabStockFromCatalogFunc(name string) models.StockListing {
	utils.InitializeLogger()
	utils.Logger.Info("grabbing stock " + name)
	stock := db.GrabCatalogStock(name)
	utils.Logger.Info("got stock " + name)
	return stock
}
