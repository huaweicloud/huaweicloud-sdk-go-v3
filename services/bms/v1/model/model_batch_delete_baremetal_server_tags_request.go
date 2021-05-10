package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteBaremetalServerTagsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchDeleteBaremetalServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteBaremetalServerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsRequest", string(data)}, " ")
}
