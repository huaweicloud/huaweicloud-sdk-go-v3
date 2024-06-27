package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashboardDataQuery struct {

	// 查询结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 分页参数，页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页参数，每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 查询开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务Id列表
	TaskIds *[]string `json:"task_ids,omitempty"`
}

func (o DashboardDataQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashboardDataQuery struct{}"
	}

	return strings.Join([]string{"DashboardDataQuery", string(data)}, " ")
}
