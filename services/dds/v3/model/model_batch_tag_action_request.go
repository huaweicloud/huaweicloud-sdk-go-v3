package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchTagActionRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *BatchOperateInstanceTagRequestBody `json:"body,omitempty"`
}

func (o BatchTagActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTagActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequest", string(data)}, " ")
}
