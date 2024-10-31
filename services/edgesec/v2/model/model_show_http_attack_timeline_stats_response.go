package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAttackTimelineStatsResponse Response Object
type ShowHttpAttackTimelineStatsResponse struct {

	// 指标类型
	StatType *string `json:"stat_type,omitempty"`

	// 分组类型
	GroupBy *string `json:"group_by,omitempty"`

	// 分组类型对应的具体的值
	GroupByValue *string `json:"group_by_value,omitempty"`

	// 时间粒度(单位：秒)
	Interval *int32 `json:"interval,omitempty"`

	// 值数组
	Values         *[]TimeStatItem `json:"values,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHttpAttackTimelineStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackTimelineStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackTimelineStatsResponse", string(data)}, " ")
}
