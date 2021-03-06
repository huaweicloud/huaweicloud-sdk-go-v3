package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *UpdateNameRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
