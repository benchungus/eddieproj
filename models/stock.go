package models

//struct for Stock object
type Stock struct {
	Name     string  `bson:"name"`
	Price    float64 `bson:"price"`
	Quantity int     `bson:"qty"`
	Beta     float64 `bson:"beta"`
}
