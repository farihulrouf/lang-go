package controller

import (
   
    "net/http"
    //"golangpro/helper"
    "encoding/json"
    "golangpro/typedata"
   // "go.mongodb.org/mongo-driver/bson"
)

func JsonHandlerTest(w http.ResponseWriter, r *http.Request) {
  // Initialize a Greeting struct
  data := &typedata.Greeting{
    Message: "I love Go!",
    Code: 100,
  }
  
  js, err := json.Marshal(data)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
  }
  w.Header().Set("Access-Control-Allow-Headers", "*")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}