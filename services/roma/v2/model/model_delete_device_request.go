package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeviceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备ID

	DeviceId int32 `json:"device_id"`
}

func (o DeleteDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceRequest", string(data)}, " ")
}
