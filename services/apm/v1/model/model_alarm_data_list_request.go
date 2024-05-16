package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmDataListRequest 告警消息请求参数。
type AlarmDataListRequest struct {

	// 页码。
	Page *int32 `json:"page,omitempty"`

	// 每页数量。
	PageSize *int32 `json:"page_size,omitempty"`

	// region英文名称。
	Region *string `json:"region,omitempty"`

	// 组件环境名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 告警状态  RECOVER：已恢复 ABNORMAL：异常 ALERT：告警中。
	Status *string `json:"status,omitempty"`

	// 告警级别  COMMON：轻微  CRITICAL：严重。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`

	// 告警开始时间。
	AlarmStartTime *string `json:"alarm_start_time,omitempty"`

	// 告警结束时间。
	AlarmEndTime *string `json:"alarm_end_time,omitempty"`

	// 采集器id。
	CollectorId *int32 `json:"collector_id,omitempty"`

	// 实例ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 环境集合。
	EnvList *[]int64 `json:"env_list,omitempty"`
}

func (o AlarmDataListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDataListRequest struct{}"
	}

	return strings.Join([]string{"AlarmDataListRequest", string(data)}, " ")
}
