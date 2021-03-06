package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTagsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTagsRequest", string(data)}, " ")
}
