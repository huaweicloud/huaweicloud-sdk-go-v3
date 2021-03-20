package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateOrDeleteInstanceTagRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteInstanceTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteInstanceTagRequest", string(data)}, " ")
}
