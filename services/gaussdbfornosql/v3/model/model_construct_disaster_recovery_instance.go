package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConstructDisasterRecoveryInstance struct {

	// 容灾实例的ID。
	Id string `json:"id"`

	// 容灾实例所在Region的code。
	RegionCode string `json:"region_code"`

	// 与当前实例建立容灾关系实例所在子网的CIDR列表。
	SubnetCidrs []string `json:"subnet_cidrs"`

	// 与当前实例建立容灾关系实例的所有节点的IP列表。
	NodeIps []string `json:"node_ips"`
}

func (o ConstructDisasterRecoveryInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstructDisasterRecoveryInstance struct{}"
	}

	return strings.Join([]string{"ConstructDisasterRecoveryInstance", string(data)}, " ")
}
