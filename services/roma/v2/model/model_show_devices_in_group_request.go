package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDevicesInGroupRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备分组ID

	GroupId int32 `json:"group_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowDevicesInGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDevicesInGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDevicesInGroupRequest", string(data)}, " ")
}
