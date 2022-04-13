package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeviceGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备分组ID

	GroupId int32 `json:"group_id"`
}

func (o ShowDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupRequest", string(data)}, " ")
}
