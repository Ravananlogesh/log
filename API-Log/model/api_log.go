package models

// This Structure is used to create Table
// Using GORM in this structure except flag all things will be inserted
type ApiCallLogStruct struct {
	
	ID            uint   `gorm:"primaryKey" json:"id"`
	URL           string `json:"url"`
	Request_Json  string `json:"request_Json"`
	Response_Json string `json:"response_Json"`
	Method        string `json:"method"`
	Source        string `json:"source"`
	Flag          string `gorm:"-" json:"flag"`
	LastId        int    `json:"lastId"`
	ErrorType     string `json:"error_type"`
}
