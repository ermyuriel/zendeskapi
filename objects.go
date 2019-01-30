package zendeskapi

import "time"

type ObjectRecordCreate struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

type ObjectTypeCreate struct {
	Key    string       `json:"key"`
	Schema ObjectSchema `json:"schema"`
}

type ObjectSchema struct {
	Properties interface{}
	Required   []string
}

type ObjectRequest struct {
	Data interface{} `json:"data"`
}

type ObjectResponse struct {
	Data struct {
		ID          string      `json:"id"`
		Type        string      `json:"type"`
		TypeVersion int         `json:"type_version"`
		Attributes  interface{} `json:"attributes"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
	} `json:"data"`
}
