package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskTimingPeriods struct {

	// 单个时间段执行的起始时间，和end_at需成对出现。计划任务类型为once时，格式为yyyy-MM-ddThh:mm:ss，其余计划任务类型时，格式为hh:mm:ss。
	BeginAt *string `json:"begin_at,omitempty"`

	// 单个时间段执行的结束时间，和begin_at需成对出现。计划任务类型为once时，格式为yyyy-MM-ddThh:mm:ss，其余计划任务类型时，格式为hh:mm:ss。
	EndAt *string `json:"end_at,omitempty"`
}

func (o TaskTimingPeriods) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTimingPeriods struct{}"
	}

	return strings.Join([]string{"TaskTimingPeriods", string(data)}, " ")
}
