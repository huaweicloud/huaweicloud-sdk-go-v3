package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddDeviceToGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备分组ID

	GroupId int32 `json:"group_id"`

	Body *BatchAddDeviceToGroupRequestBody `json:"body,omitempty"`
}

func (o BatchAddDeviceToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeviceToGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDeviceToGroupRequest", string(data)}, " ")
}
