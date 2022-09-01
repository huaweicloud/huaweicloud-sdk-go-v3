package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrecheckDisasterRecoveryInstance struct {

	// 与当前实例建立容灾关系实例的vpc网段。
	VpcCidr string `json:"vpc_cidr" xml:"vpc_cidr"`

	// 与当前实例建立容灾关系实例的规格码。
	SpecCode string `json:"spec_code" xml:"spec_code"`

	// 与当前实例建立容灾关系实例的节点IP列表。
	NodeIps []string `json:"node_ips" xml:"node_ips"`
}

func (o PrecheckDisasterRecoveryInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckDisasterRecoveryInstance struct{}"
	}

	return strings.Join([]string{"PrecheckDisasterRecoveryInstance", string(data)}, " ")
}
