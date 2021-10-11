package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UploadProtocolMappingsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadProtocolMappingsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadProtocolMappingsResponse struct{}"
	}

	return strings.Join([]string{"UploadProtocolMappingsResponse", string(data)}, " ")
}
