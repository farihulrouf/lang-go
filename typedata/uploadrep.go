package typedata
import (
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)
type Uploadrep struct {
	Exchange      string   `bson:"exchange" json:"exchange"`
	Sell		  int      `bson:"sell" json:"sell"`
}
