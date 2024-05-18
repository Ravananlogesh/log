package repositories

import (
	"api-log/helpers"
	models "api-log/model"

	"gorm.io/gorm"
)

type ApiRepoStruct struct {
	db *gorm.DB
}

func ApiCallRepo(db *gorm.DB) *ApiRepoStruct {
	return &ApiRepoStruct{db: db}
}

func (r *ApiRepoStruct) Save(lApiCallLog *models.ApiCallLogStruct) error {
	return r.db.Create(lApiCallLog).Error
}
func (r *ApiRepoStruct) Update(lApiCallLog *models.ApiCallLogStruct) error {
	return r.db.Save(lApiCallLog).Error
}

func (r *ApiRepoStruct) Get(lastId int) (*models.ApiCallLogStruct, error) {
	var lApiCallLog models.ApiCallLogStruct
	result := r.db.First(&lApiCallLog, "last_id=?", lastId)
	if result.Error != nil {
		return nil, helpers.ErrReturn(result.Error)
	}
	return &lApiCallLog, nil
}
