package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GovPolicyDetailPolicies struct {

	// 治理策略ID。
	Id *string `json:"id,omitempty"`

	// 治理策略名称。
	Name *string `json:"name,omitempty"`

	// 治理类型，支持填写retry、rate-limiting、loadbalance、circuit-breaker、instance-isolation、fault-injection和bulkhead。
	Kind *string `json:"kind,omitempty"`

	// 启用状态，支持enabled和disabled。
	Status *string `json:"status,omitempty"`

	Selector *GovSelector `json:"selector,omitempty"`

	// 治理策略定义内容。
	Spec *interface{} `json:"spec,omitempty"`
}

func (o GovPolicyDetailPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GovPolicyDetailPolicies struct{}"
	}

	return strings.Join([]string{"GovPolicyDetailPolicies", string(data)}, " ")
}
