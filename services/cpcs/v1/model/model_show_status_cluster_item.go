package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatusClusterItem struct {

	// 成功率
	SuccessRate *float64 `json:"success_rate,omitempty"`

	// 失败率
	FailRate *float64 `json:"fail_rate,omitempty"`

	// 成功次数
	SuccessCount *int64 `json:"success_count,omitempty"`

	// 总次数
	TotalCount *int64 `json:"total_count,omitempty"`

	// 平均值
	AverageValue *float64 `json:"average_value,omitempty"`

	// 总数
	TotalValue *float64 `json:"total_value,omitempty"`

	// 指标是否超出限制
	Limit *bool `json:"limit,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o ShowStatusClusterItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusClusterItem struct{}"
	}

	return strings.Join([]string{"ShowStatusClusterItem", string(data)}, " ")
}
