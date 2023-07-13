package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceRequest Request Object
type DeleteDeviceRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 终端设备ID
	DeviceId string `json:"device_id"`
}

func (o DeleteDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceRequest", string(data)}, " ")
}
