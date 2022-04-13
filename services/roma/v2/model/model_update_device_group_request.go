package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeviceGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备分组ID

	GroupId int32 `json:"group_id"`

	Body *UpdateDeviceGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupRequest", string(data)}, " ")
}
