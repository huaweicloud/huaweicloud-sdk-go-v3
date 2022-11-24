package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询趋势图入参。
type TrendParam struct {
	ViewConfig *TrendView `json:"view_config,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o TrendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendParam struct{}"
	}

	return strings.Join([]string{"TrendParam", string(data)}, " ")
}
