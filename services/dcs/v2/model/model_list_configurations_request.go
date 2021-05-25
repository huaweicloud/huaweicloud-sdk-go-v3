package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConfigurationsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ListConfigurationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}
