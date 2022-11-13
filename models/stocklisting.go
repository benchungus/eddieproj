package models

//struct for Stock object
type StockListing struct {
	Name  string  `bson:"name"`
	Price float64 `bson:"price"`
	Beta  float64 `bson:"beta"`
}
