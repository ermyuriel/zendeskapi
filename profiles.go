package zendeskapi

type ProfileRequest struct {
	Data Profile `json:"profile"`
}

type Profile struct {
	Source      string      `json:"source"`
	Type        string      `json:"type"`
	Identifiers interface{} `json:"identifiers"`
	Attributes  interface{} `json:"attributes"`
}
