package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH服务可用区对象
type AvailabilityZones struct {

	// 可用区id
	Id string `json:"id"`

	// 可用区显示名称
	DisplayName string `json:"display_name"`

	// region id
	RegionId string `json:"region_id"`

	// 可用区状态
	Status string `json:"status"`

	// 可用区类型
	Type *string `json:"type,omitempty"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}
