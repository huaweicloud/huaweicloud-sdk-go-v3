package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeviceTwinRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 设备ID
	DeviceId string `json:"device_id" xml:"device_id"`

	Body *TwinUpdateDetail `json:"body,omitempty" xml:"body"`
}

func (o UpdateDeviceTwinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTwinRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTwinRequest", string(data)}, " ")
}
