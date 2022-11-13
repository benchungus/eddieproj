package db

import (
	"context"
	"log"

	"github.com/benchungus/api/models"
	"github.com/benchungus/api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// searches for all stocks in riskanal db and returns as an array
func GrabAllStocks() []models.Stock {
	utils.InitializeLogger()
	client := ConnDB()
	stockCollection := client.Database("riskanal").Collection("stocks")

	//generating cursor for this search
	cursor, err := stockCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var stocks []models.Stock

	//adding all found stocks to an array
	if err = cursor.All(context.TODO(), &stocks); err != nil {
		log.Fatal(err)
	}

	return stocks
}

// inserts new stock into riskanal db
func InsertStock(newStock models.Stock) {
	utils.InitializeLogger()
	client := ConnDB()
	stockCollection := client.Database("riskanal").Collection("stocks")
	_, err := stockCollection.InsertOne(context.TODO(), newStock)
	if err != nil {
		log.Panic(err)
	}
}

func GrabNamedStock(name string) models.Stock {
	utils.InitializeLogger()
	client := ConnDB()
	var stock models.Stock
	stockCollection := client.Database("riskanal").Collection("stocks")

	//find and decode into stock
	err := stockCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&stock)
	if err != nil {
		log.Panic(err)
	}
	return stock
}

func GrabBond(name string) models.Stock {
	utils.InitializeLogger()
	client := ConnDB()
	var stock models.Stock
	stockCollection := client.Database("riskanal").Collection("bonds")

	//find and decode into stock
	err := stockCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&stock)
	if err != nil {
		log.Panic(err)
	}
	return stock
}

func GrabCatalogStock(name string) models.StockListing {
	utils.InitializeLogger()
	client := ConnDB()
	var stock models.StockListing
	stockCollection := client.Database("riskanal").Collection("stockcatalog")

	//find and decode into stock
	err := stockCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&stock)
	if err != nil {
		log.Panic(err)
	}
	return stock
}

// updates a stocks's quantity by searching for name
func UpdateStock(name string, quantity int) {
	utils.InitializeLogger()
	client := ConnDB()
	stockCollection := client.Database("riskanal").Collection("stocks")
	_, err := stockCollection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.D{{"qty", quantity}}}})
	if err != nil {
		log.Panic(err)
	}

}

func UpdateBond(name string, quantity int) {
	utils.InitializeLogger()
	client := ConnDB()
	stockCollection := client.Database("riskanal").Collection("bonds")
	_, err := stockCollection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.D{{"qty", quantity}}}})
	if err != nil {
		log.Panic(err)
	}

}

// deletes a stock from the database
func DeleteStock(name string) {
	utils.InitializeLogger()
	client := ConnDB()
	stockCollection := client.Database("riskanal").Collection("stocks")
	_, err := stockCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	if err != nil {
		log.Panic(err)
	}
}
