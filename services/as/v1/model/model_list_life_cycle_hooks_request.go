package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLifeCycleHooksRequest struct {
	// 伸缩组标识。

	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ListLifeCycleHooksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLifeCycleHooksRequest struct{}"
	}

	return strings.Join([]string{"ListLifeCycleHooksRequest", string(data)}, " ")
}
