package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMqsInstanceRequest Request Object
type ShowMqsInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowMqsInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceRequest", string(data)}, " ")
}
