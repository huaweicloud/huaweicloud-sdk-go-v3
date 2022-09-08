package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控项数据结构
type MonitorItem struct {

	// 监控项id
	Id *int64 `json:"id,omitempty"`

	// 环境id
	EnvId *int64 `json:"env_id,omitempty"`

	// 采集器id
	CollectorId *int32 `json:"collector_id,omitempty"`

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty"`

	// 采集器展示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 采集间隔
	CollectInterval *int32 `json:"collect_interval,omitempty"`

	// 是否禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 修改采集状态用户id
	StatusChangeUserId *string `json:"status_change_user_id,omitempty"`

	// 修改采集状态用户名称
	StatusChangeUserName *string `json:"status_change_user_name,omitempty"`

	// 修改采集状态时间
	StatusChangeTime *string `json:"status_change_time,omitempty"`

	// 修改采集配置用户id
	ConfigChangeUserId *string `json:"config_change_user_id,omitempty"`

	// 修改采集配置用户名称
	ConfigChangeUserName *string `json:"config_change_user_name,omitempty"`

	// 修改采集配置时间
	ConfigChangeTime *string `json:"config_change_time,omitempty"`
}

func (o MonitorItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorItem struct{}"
	}

	return strings.Join([]string{"MonitorItem", string(data)}, " ")
}
