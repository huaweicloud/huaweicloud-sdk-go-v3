package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveMonitorItemParam struct {

	// 监控项id。
	MonitorItemId int64 `json:"monitor_item_id"`

	// 采集间隔。
	Interval *int32 `json:"interval,omitempty"`

	// 环境id。
	EnvId int32 `json:"env_id"`

	// 配置项列表。
	ConfigValueList *[]ConfigItem `json:"config_value_list,omitempty"`
}

func (o SaveMonitorItemParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveMonitorItemParam struct{}"
	}

	return strings.Join([]string{"SaveMonitorItemParam", string(data)}, " ")
}
