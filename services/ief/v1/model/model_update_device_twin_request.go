package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeviceTwinRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 设备ID

	DeviceId string `json:"device_id"`

	Body *TwinUpdateDetail `json:"body,omitempty"`
}

func (o UpdateDeviceTwinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTwinRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTwinRequest", string(data)}, " ")
}
