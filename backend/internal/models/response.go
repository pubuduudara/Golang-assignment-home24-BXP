package models

// common response format
type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"` // can be actual result or error message
}

