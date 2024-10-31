package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAttackTopStatsResponse Response Object
type ShowHttpAttackTopStatsResponse struct {

	// 指标类型
	StatType *string `json:"stat_type,omitempty"`

	// 分组类型
	GroupBy *string `json:"group_by,omitempty"`

	Values *[]CommonStatItem `json:"values,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime        *int64 `json:"end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHttpAttackTopStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackTopStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackTopStatsResponse", string(data)}, " ")
}
