package typedata
import (
	//"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Partner struct {
	Code      string    	        `bson:"code" json:"code"`
	Createby  string    		    `bson:"createby" json:"createby"`
	Createon  primitive.DateTime	`bson:"createon" json:"createon"`
	Enable	  bool					`bson:"enable" json:"enable"`
	Name	  string		        `bson:"name" json:"name"`
	Updateby  string	            `bson:"updateby" json:"updateby"`
	Updateon  primitive.DateTime 	`bson:"updateon" json:"updateon"`
}
