package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowClusterRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowClusterRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterRequest", string(data)}, " ")
}
