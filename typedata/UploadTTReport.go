package typedata
import (
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)
type UploadTTReport struct {
	Exchange          string    	        `bson:"exchange" json:"exchange"`
	Transactiontype   string    		    `bson:"transactiontype" json:"transactiontype"`
	Fullname          string	            `bson:"fullname" json:"fullname"`
	Unrealize	      bool					`bson:"unrealize" json:"unrealize"`
	Filename	      string		        `bson:"filename" json:"filename"`
	Contractcode      string	                `bson:"contractcode" json:"contractcode"`
	Currency          string 	            `bson:"currency" json:"currency"`
}
