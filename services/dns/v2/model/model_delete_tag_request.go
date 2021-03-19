package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTagRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Key string `json:"key"`
}

func (o DeleteTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
