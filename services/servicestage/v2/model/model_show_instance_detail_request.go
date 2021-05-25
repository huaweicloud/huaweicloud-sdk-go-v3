package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceDetailRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailRequest", string(data)}, " ")
}
