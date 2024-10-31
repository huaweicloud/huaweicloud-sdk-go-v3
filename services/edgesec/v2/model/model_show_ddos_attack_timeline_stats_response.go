package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdosAttackTimelineStatsResponse Response Object
type ShowDdosAttackTimelineStatsResponse struct {

	// 指标类型
	StatType *string `json:"stat_type,omitempty"`

	// 分组类型
	GroupBy *string `json:"group_by,omitempty"`

	// 时间粒度(单位：秒)
	Interval *int32 `json:"interval,omitempty"`

	// 值数组
	Values         *[]TimeStatItem `json:"values,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowDdosAttackTimelineStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdosAttackTimelineStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDdosAttackTimelineStatsResponse", string(data)}, " ")
}
