package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeviceGroupRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 设备分组ID
	GroupId int32 `json:"group_id" xml:"group_id"`
}

func (o DeleteDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupRequest", string(data)}, " ")
}
