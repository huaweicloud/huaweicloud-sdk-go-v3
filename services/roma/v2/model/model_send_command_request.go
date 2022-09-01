package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendCommandRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id" xml:"device_id"`

	Body *SendCommandRequestBody `json:"body,omitempty" xml:"body"`
}

func (o SendCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendCommandRequest struct{}"
	}

	return strings.Join([]string{"SendCommandRequest", string(data)}, " ")
}
