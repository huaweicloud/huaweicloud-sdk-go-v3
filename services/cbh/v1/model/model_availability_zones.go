package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云堡垒机服务可用分区信息。
type AvailabilityZones struct {

	// 云堡垒机服务可用区ID。
	Id string `json:"id"`

	// 云堡垒机服务可用分区显示名称。
	DisplayName string `json:"display_name"`

	// 云堡垒机服务分区ID。
	RegionId string `json:"region_id"`

	// 云堡垒机服务可用区状态。
	Status string `json:"status"`

	// 云堡垒机服务可用区类型。
	Type *string `json:"type,omitempty"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}
