package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZones 云堡垒机服务可用分区信息。
type AvailabilityZones struct {

	// 可用区ID。
	RegionId string `json:"region_id"`

	// 可用分区ID。
	Id string `json:"id"`

	// 可用分区显示名称。
	DisplayName string `json:"display_name"`

	// 可用分区状态。 - Running：运行中
	Status string `json:"status"`

	// 可用分区类型。 - Core：核心可用区 - Dedicated：专属可用区，只对内部客户开放
	Type *string `json:"type,omitempty"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}
