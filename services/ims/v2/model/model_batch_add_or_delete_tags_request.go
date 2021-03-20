package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddOrDeleteTagsRequest struct {
	ImageId string `json:"image_id"`

	Body *BatchAddOrDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddOrDeleteTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsRequest", string(data)}, " ")
}
