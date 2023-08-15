package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceRequest Request Object
type UpdateDeviceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id"`

	Body *UpdateDeviceRequestBody `json:"body,omitempty"`
}

func (o UpdateDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequest", string(data)}, " ")
}
