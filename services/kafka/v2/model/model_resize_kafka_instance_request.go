package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeKafkaInstanceRequest Request Object
type ResizeKafkaInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeEngineInstanceReq `json:"body,omitempty"`
}

func (o ResizeKafkaInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeKafkaInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeKafkaInstanceRequest", string(data)}, " ")
}
