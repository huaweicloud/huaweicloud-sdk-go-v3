package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAttackDistributionStatsResponse Response Object
type ShowHttpAttackDistributionStatsResponse struct {

	// 安全统计指标类型。
	StatType *string `json:"stat_type,omitempty"`

	// 分组统计维度
	GroupBy *string `json:"group_by,omitempty"`

	// 统计数量
	StatNum *int64 `json:"stat_num,omitempty"`

	// 统计数组
	Values *[]CommonStatItem `json:"values,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime        *int64 `json:"end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHttpAttackDistributionStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackDistributionStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackDistributionStatsResponse", string(data)}, " ")
}
