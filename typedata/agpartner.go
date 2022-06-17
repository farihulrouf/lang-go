package typedata
import (
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)
type AgaPrtner struct {
	Id          Dcode    	         `bson:"_id" json:"_id"`
	Total 	    int		             `bson:"total" json:"total"`
}

type Dcode struct {
	Code  string `bson:"code" json:"code"`
	//Code  string `bson:"code" json:"code"`
}


