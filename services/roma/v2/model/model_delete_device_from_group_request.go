package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeviceFromGroupRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id" xml:"device_id"`

	// 设备分组ID
	GroupId int32 `json:"group_id" xml:"group_id"`
}

func (o DeleteDeviceFromGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceFromGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceFromGroupRequest", string(data)}, " ")
}
