package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineSchedule struct {

	// 任务ID
	Uuid *string `json:"uuid,omitempty"`

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 是否可用
	Enable *string `json:"enable,omitempty"`

	// 一周内具体时间
	DaysOfWeek *[]int32 `json:"days_of_week,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o PipelineSchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSchedule struct{}"
	}

	return strings.Join([]string{"PipelineSchedule", string(data)}, " ")
}
