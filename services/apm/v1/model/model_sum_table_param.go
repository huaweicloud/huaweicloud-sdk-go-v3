package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 汇聚数据表格传参。
type SumTableParam struct {

	// 上次请求的id。
	RequestId *string `json:"request_id,omitempty"`

	ViewConfig *SumTableView `json:"view_config"`

	// 策略。
	Strategy *string `json:"strategy,omitempty"`

	// 当前页码。
	Page int32 `json:"page"`

	// 每页数据总数。
	PageSize int32 `json:"page_size"`

	// 排序规则。
	OrderBy *string `json:"order_by,omitempty"`

	// 搜索关键字。
	SearchWord *string `json:"search_word,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 环境id。
	EnvId int64 `json:"env_id"`

	// 起始时间。
	StartTime string `json:"start_time"`

	// 结束时间。
	EndTime string `json:"end_time"`
}

func (o SumTableParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SumTableParam struct{}"
	}

	return strings.Join([]string{"SumTableParam", string(data)}, " ")
}
