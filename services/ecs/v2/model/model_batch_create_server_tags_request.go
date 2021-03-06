package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateServerTagsRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`

	Body *BatchCreateServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateServerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsRequest", string(data)}, " ")
}
