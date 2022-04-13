package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeviceTwinRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 终端设备ID

	DeviceId string `json:"device_id"`
}

func (o ShowDeviceTwinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTwinRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceTwinRequest", string(data)}, " ")
}
