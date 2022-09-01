package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控项数据结构
type MonitorItem struct {

	// 监控项id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 环境id
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 采集器id
	CollectorId *int32 `json:"collector_id,omitempty" xml:"collector_id"`

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty" xml:"collector_name"`

	// 采集器展示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 采集间隔
	CollectInterval *int32 `json:"collect_interval,omitempty" xml:"collect_interval"`

	// 是否禁用
	Disabled *bool `json:"disabled,omitempty" xml:"disabled"`

	// 修改采集状态用户id
	StatusChangeUserId *string `json:"status_change_user_id,omitempty" xml:"status_change_user_id"`

	// 修改采集状态用户名称
	StatusChangeUserName *string `json:"status_change_user_name,omitempty" xml:"status_change_user_name"`

	// 修改采集状态时间
	StatusChangeTime *string `json:"status_change_time,omitempty" xml:"status_change_time"`

	// 修改采集配置用户id
	ConfigChangeUserId *string `json:"config_change_user_id,omitempty" xml:"config_change_user_id"`

	// 修改采集配置用户名称
	ConfigChangeUserName *string `json:"config_change_user_name,omitempty" xml:"config_change_user_name"`

	// 修改采集配置时间
	ConfigChangeTime *string `json:"config_change_time,omitempty" xml:"config_change_time"`
}

func (o MonitorItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorItem struct{}"
	}

	return strings.Join([]string{"MonitorItem", string(data)}, " ")
}
