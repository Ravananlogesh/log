package main

import (
	"api-log/helpers"
	models "api-log/model"
	"log"

	"api-log/util"
)

func main() {

	lDebug := new(helpers.HelperStruct)
	lDebug.Init()
	lInput := models.ApiCallLogStruct{
		URL:           "www.google.com",
		Request_Json:  `{"logesh":123}`,
		Response_Json: `{"logesh":321}`,
		Method:        "GET",
		Source:        "any",
		Flag:          "INSERT",
		LastId:        1,
		ErrorType:     "N",
	}

	lLastId, err := util.ApiLogEntry(lDebug, lInput)
	if err != nil {
		log.Fatalf("ApiLogEntry failed: %v", err)
	} else {
		log.Println("last Id is ===>", lLastId)
    }
}
