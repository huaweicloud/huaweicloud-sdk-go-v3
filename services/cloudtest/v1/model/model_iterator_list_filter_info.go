package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IteratorListFilterInfo 过滤条件
type IteratorListFilterInfo struct {

	// pi过滤条件
	PiSprints *[]IssueListPiFilterInfo `json:"pi_sprints,omitempty"`

	// 计划结束间过滤开始时间点
	PlanEndDateStart *sdktime.SdkTime `json:"plan_end_date_start,omitempty"`

	// 计划结束时间过滤结束时间点
	PlanEndDateEnd *sdktime.SdkTime `json:"plan_end_date_end,omitempty"`
}

func (o IteratorListFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IteratorListFilterInfo struct{}"
	}

	return strings.Join([]string{"IteratorListFilterInfo", string(data)}, " ")
}
