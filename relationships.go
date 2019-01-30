package zendeskapi

type RelationshipRequest struct {
	Data interface{} `json:"data"`
}

type RelationshipTypeCreate struct {
	Key    string      `json:"key"`
	Source interface{} `json:"source"`
	Target interface{} `json:"target"`
}

type RelationshipRecordCreate struct {
	RelationshipType string      `json:"relationship_type"`
	Source           interface{} `json:"source"`
	Target           interface{} `json:"target"`
}

type RelationshipSearchResponse struct {
	Data  []Relationship `json:"data"`
	Links struct {
		Previous interface{} `json:"previous"`
		Next     interface{} `json:"next"`
	} `json:"links"`
}

type Relationship struct {
	ID     string `json:"id"`
	Target string `json:"target"`
	Ref    string `json:"ref"`
}
