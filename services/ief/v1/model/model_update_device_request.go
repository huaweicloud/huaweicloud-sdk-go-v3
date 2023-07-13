package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceRequest Request Object
type UpdateDeviceRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 设备ID
	DeviceId string `json:"device_id"`

	Body *EdgemgrDevicesUpdate `json:"body,omitempty"`
}

func (o UpdateDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequest", string(data)}, " ")
}
