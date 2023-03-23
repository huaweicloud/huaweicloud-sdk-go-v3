package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeEngineInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine string `json:"engine"`

	Body *ResizeEngineInstanceReq `json:"body,omitempty"`
}

func (o ResizeEngineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceRequest", string(data)}, " ")
}
