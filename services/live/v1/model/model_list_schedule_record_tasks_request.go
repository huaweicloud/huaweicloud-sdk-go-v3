package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleRecordTasksRequest Request Object
type ListScheduleRecordTasksRequest struct {

	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（任务开始结束时间必须在查询时间范围内）。
	EndTime int64 `json:"end_time"`

	// 推流域名
	Domain *string `json:"domain,omitempty"`

	// 应用名称，在domain有值的情况下生效
	App *string `json:"app,omitempty"`

	// 流名称，在domain和app有值的情况下生效
	Stream *string `json:"stream,omitempty"`

	// 录制任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 分页场景下的偏移量，表示从此偏移量开始查询，offset大于等于0，缺省值为0，需要每次查询增加limit的值才能查询全部数据
	Offset *int32 `json:"offset,omitempty"`

	// 分页场景下的每页记录数，取值范围[1,100]，缺省值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduleRecordTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleRecordTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleRecordTasksRequest", string(data)}, " ")
}
