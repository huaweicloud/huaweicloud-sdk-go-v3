package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RawTableParam 获取原始数据表格入参。
type RawTableParam struct {

	// 上一次扫描的数据ID。
	LastRowId *string `json:"last_row_id,omitempty"`

	ViewConfig *RawTableView `json:"view_config"`

	// 当前页码。
	Page int32 `json:"page"`

	// 每页数据总数。
	PageSize int32 `json:"page_size"`

	// 排序。
	OrderBy *string `json:"order_by,omitempty"`

	// 搜索关键字。
	SearchWord *string `json:"search_word,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 环境id。
	EnvId int64 `json:"env_id"`

	// 开始时间。
	StartTime string `json:"start_time"`

	// 结束时间。
	EndTime string `json:"end_time"`
}

func (o RawTableParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RawTableParam struct{}"
	}

	return strings.Join([]string{"RawTableParam", string(data)}, " ")
}
