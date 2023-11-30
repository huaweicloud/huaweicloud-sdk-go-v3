package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceRequest Request Object
type ResizeInstanceRequest struct {

	// 消息引擎的类型。支持的类型为rocketmq。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeEngineInstanceReq `json:"body,omitempty"`
}

func (o ResizeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequest", string(data)}, " ")
}
