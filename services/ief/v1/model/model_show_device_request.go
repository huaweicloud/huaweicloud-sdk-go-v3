package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeviceRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 终端设备ID
	DeviceId string `json:"device_id" xml:"device_id"`
}

func (o ShowDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceRequest", string(data)}, " ")
}
