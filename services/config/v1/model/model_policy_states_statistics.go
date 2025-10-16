package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyStatesStatistics 合规统计结果
type PolicyStatesStatistics struct {

	// 资源总数
	TotalResourceCount *int32 `json:"total_resource_count,omitempty"`

	// 不合规资源数量
	NonCompliantResourceCount *int32 `json:"non_compliant_resource_count,omitempty"`

	// 合规规则总数
	TotalPolicyCount *int32 `json:"total_policy_count,omitempty"`

	// 不合规规则数量
	NonCompliantPolicyCount *int32 `json:"non_compliant_policy_count,omitempty"`

	// 统计日期
	StatisticDate *string `json:"statistic_date,omitempty"`
}

func (o PolicyStatesStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatesStatistics struct{}"
	}

	return strings.Join([]string{"PolicyStatesStatistics", string(data)}, " ")
}
