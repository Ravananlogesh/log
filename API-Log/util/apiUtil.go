package util

import (
	"api-log/config"
	"api-log/helpers"
	models "api-log/model"
	"api-log/repositories"
	"fmt"
)

/*
Pupose: This method is used to store the data for endppoint datatable Using GORM
Parameters:

	send ApiCallLog as a parameter to this method

Response:

	    *On Sucess
	    =========
	    In case of a successful execution of this method, you will get the Error Message as nil data
		from Return value

	    !On Error
	    ========
	    In case of any exception during the execution of this method you will get the
	    error details. the calling program should handle the error

Author: LogeshKumar P
Date: 17 May 2024
*/
func ApiLogEntry(pDebug *helpers.HelperStruct, pInput models.ApiCallLogStruct) (uint, error) {
	pDebug.Log(helpers.Statement, "ApiLogEntry (+)")

	// create a instace to hold last inserted Id
	var lLastId uint
	// Calling LocalDbConect method in config to estabish the database connection
	db, lErr := config.LocalDbConnect()
	if lErr != nil {
		return lLastId, helpers.ErrReturn(lErr)
	}

	// Auto migrate the schema to auto create table based on structure and based on value create column
	lCreateErr := db.AutoMigrate(&models.ApiCallLogStruct{})
	if lCreateErr != nil {
		return lLastId, helpers.ErrReturn(lCreateErr)
	}

	repo := repositories.ApiCallRepo(db)

	if pInput.Flag == "INSERT" {
		lErr := repo.Save(&pInput)
		if lErr != nil {
			return lLastId, helpers.ErrReturn(lErr)
		}
		lLastId = pInput.ID
	} else if pInput.Flag == "UPDATE" {

		//Updating the log

		pInput.Method = "POST"
		updateErr := repo.Update(&pInput)
		if updateErr != nil {
			return lLastId, helpers.ErrReturn(updateErr)
		}
	}
	// Fetching the same log
	fetchedLog, lErr := repo.Get(pInput.LastId)
	if lErr != nil {
		return lLastId, helpers.ErrReturn(lErr)
	}
	fmt.Printf("Fetched Log: %+v\n", fetchedLog)

	pDebug.Log(helpers.Statement, "ApiLogEntry (-)")
	return lLastId, nil
}
