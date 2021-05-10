package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateBaremetalServerTagsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchCreateBaremetalServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateBaremetalServerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateBaremetalServerTagsRequest", string(data)}, " ")
}
