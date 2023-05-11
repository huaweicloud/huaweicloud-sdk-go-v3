package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息
type InstanceInfo struct {

	// 虚拟机名称
	Name string `json:"name"`

	// 虚拟机Id
	Id string `json:"id"`

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 虚拟机规格
	FlavorId string `json:"flavor_id"`

	// 虚拟机状态
	Status string `json:"status"`

	// 销售模型（spot/onDemand）
	SellMode string `json:"sell_mode"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}
