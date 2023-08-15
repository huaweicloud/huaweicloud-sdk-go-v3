package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceRequest Request Object
type ShowDeviceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id"`
}

func (o ShowDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceRequest", string(data)}, " ")
}
