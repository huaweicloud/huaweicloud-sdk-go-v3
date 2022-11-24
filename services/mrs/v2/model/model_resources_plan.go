package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesPlan struct {

	// 资源计划的周期类型，当前只允许以下类型：  daily
	PeriodType string `json:"period_type"`

	// 资源计划的起始时间，格式为“hour:minute”，表示时间在0:00-23:59之间。
	StartTime string `json:"start_time"`

	// 资源计划的结束时间，格式与“start_time”相同，不早于start_time表示的时间，且与start_time间隔不小于30min。
	EndTime string `json:"end_time"`

	// 资源计划内该节点组的最小保留节点数。 取值范围：[0～500]
	MinCapacity int32 `json:"min_capacity"`

	// 资源计划内该节点组的最大保留节点数。 取值范围：[0～500]
	MaxCapacity int32 `json:"max_capacity"`
}

func (o ResourcesPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesPlan struct{}"
	}

	return strings.Join([]string{"ResourcesPlan", string(data)}, " ")
}
