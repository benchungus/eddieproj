package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/benchungus/api/db"
	"github.com/benchungus/api/models"

	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

// handles GET request from /stocks endpoint, returns all stocks in db
func GetStocks(c *gin.Context) {
	utils.InitializeLogger()
	utils.Logger.Info("grabbing all stocks")
	stocks := db.GrabAllStocks()
	utils.Logger.Info("got all stocks")
	c.IndentedJSON(http.StatusOK, stocks)
}

func GetBond(c *gin.Context) {
	name := "bond"
	utils.InitializeLogger()
	utils.Logger.Info("grabbing " + name)
	stock := db.GrabBond(name)
	utils.Logger.Info("got " + name)
	c.IndentedJSON(http.StatusOK, stock)
}

// handles POST request from /stocks endpoint, inserts new stock
func NewStock(c *gin.Context) {
	utils.InitializeLogger()
	name := c.Param("name")
	stockCat := GrabStockFromCatalogFunc(name)
	var newStock models.Stock
	newStock.Beta = float64(stockCat.Beta)
	newStock.Name = name
	newStock.Price = float64(stockCat.Price)
	qty := c.Param("qty")
	newStock.Quantity, _ = strconv.Atoi(qty)
	utils.Logger.Info("inserting stock " + newStock.Name)
	db.InsertStock(newStock)
	utils.Logger.Info("inserted stock named " + newStock.Name)
}

// handles POST request from /stocks/:name endpoint, updates qty of stock
func ChangeQty(c *gin.Context) {
	utils.InitializeLogger()
	name := c.Param("name")
	var newStock models.Stock
	err := json.NewDecoder(c.Request.Body).Decode(&newStock)
	if err != nil {
		log.Panic(err)
	}
	utils.Logger.Info("updating stock " + name)
	db.UpdateStock(name, newStock.Quantity)
	utils.Logger.Info("updated stock " + newStock.Name)
}

func ChangeBond(c *gin.Context) {
	utils.InitializeLogger()
	qty := c.Param("qty")
	name := "bond"
	var newStock models.Stock
	stockCat := GrabStockFromCatalogFunc(name)
	newStock.Beta = float64(stockCat.Beta)
	newStock.Name = name
	newStock.Price = float64(stockCat.Price)
	newStock.Quantity, _ = strconv.Atoi(qty)
	err := json.NewDecoder(c.Request.Body).Decode(&newStock)
	if err != nil {
		log.Panic(err)
	}
	utils.Logger.Info("updating " + name)
	db.UpdateBond(name, newStock.Quantity)
	utils.Logger.Info("updated " + newStock.Name)
}

func GetStockByName(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("grabbing stock " + name)
	stock := db.GrabNamedStock(name)
	utils.Logger.Info("got stock " + name)
	c.IndentedJSON(http.StatusOK, stock)
}

// handles DELETE request from /stocks/:name endpoint, deletes said stock
func DeleteStock(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("deleting stock named " + name)
	db.DeleteStock(name)
	utils.Logger.Info("deleted stock named " + name)
}

func AggBeta(c *gin.Context) {
	totBet := 0.0
	totMon := 0.0
	utils.InitializeLogger()
	utils.Logger.Info("grabbing all stocks")
	stocks := db.GrabAllStocks()
	bond := db.GrabBond("bond")
	utils.Logger.Info("got all stocks")
	for i := 0; i < len(stocks); i++ {
		totBet += stocks[i].Beta * stocks[i].Price * float64(stocks[i].Quantity)
		totMon += stocks[i].Price * float64(stocks[i].Quantity)
	}
	totBet += bond.Beta * bond.Price * float64(bond.Quantity)
	totMon += bond.Price * float64(bond.Quantity)
	ans := totBet / totMon
	c.IndentedJSON(http.StatusOK, ans)
}
