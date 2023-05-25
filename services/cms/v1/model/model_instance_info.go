package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 虚拟机列表
type InstanceInfo struct {

	// 弹性云服务器名称。
	Name string `json:"name"`

	// 弹性云服务器id。
	Id string `json:"id"`

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 弹性云服务器规格
	FlavorId string `json:"flavor_id"`

	// 弹性云服务器状态
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
