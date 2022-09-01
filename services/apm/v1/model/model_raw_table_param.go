package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取原始数据表格入参
type RawTableParam struct {

	// 上一次扫描的数据ID
	LastRowId *string `json:"lastRowId,omitempty" xml:"lastRowId"`

	ViewConfig *RawTableView `json:"view_config,omitempty" xml:"view_config"`

	// 当前页码
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 每页数据总数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 排序
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// 搜索关键字
	SearchWord *string `json:"search_word,omitempty" xml:"search_word"`

	// 实例id
	InstanceId *int64 `json:"instance_id,omitempty" xml:"instance_id"`

	// 监控项id
	MonitorItemId *int64 `json:"monitor_item_id,omitempty" xml:"monitor_item_id"`

	// 环境id
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o RawTableParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RawTableParam struct{}"
	}

	return strings.Join([]string{"RawTableParam", string(data)}, " ")
}
