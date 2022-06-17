package typedata
import (
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)
type Selling struct {	
	Exchange   string `bson:"exchange" json:"exchange"`
	Sell       float32 `bson:"sell" json:"sell"`
}
