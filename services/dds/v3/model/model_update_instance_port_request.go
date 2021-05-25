package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstancePortRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *UpdatePortRequestBody `json:"body,omitempty"`
}

func (o UpdateInstancePortRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstancePortRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstancePortRequest", string(data)}, " ")
}
