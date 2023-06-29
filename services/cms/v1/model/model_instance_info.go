package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceInfo 实例列表
type InstanceInfo struct {

	// 实例名称。
	Name string `json:"name"`

	// 实例ID。
	Id string `json:"id"`

	// 可用区ID
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 实例规格
	FlavorId string `json:"flavor_id"`

	// 实例状态
	Status string `json:"status"`

	// 销售模型，枚举值 spot：竞价实例 onDemand：按需实例
	SellMode string `json:"sell_mode"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}
