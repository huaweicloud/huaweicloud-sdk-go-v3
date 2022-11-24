package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警信息视图。
type AlarmDataVo struct {

	// 告警通知id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// region中事件的id。
	RegionAlarmEventId *int64 `json:"region_alarm_event_id,omitempty"`

	// 应用名称。
	BusinessName *string `json:"business_name,omitempty"`

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 版本。
	VersionNumber *int32 `json:"version_number,omitempty"`

	// 告警规则类别。
	AlarmRuleType *string `json:"alarm_rule_type,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 处理单元。
	ProcessUnit *string `json:"process_unit,omitempty"`

	// 区域名称。
	Region *string `json:"region,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 实例ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 模板id。
	TemplateId *int64 `json:"template_id,omitempty"`

	// 告警规则id。
	AlarmRuleId *int64 `json:"alarm_rule_id,omitempty"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 采集器id。
	CollectorId *int32 `json:"collector_id,omitempty"`

	// 采集器名称。
	CollectorName *string `json:"collector_name,omitempty"`

	// 告警规则名称。
	AlarmRuleName *string `json:"alarm_rule_name,omitempty"`

	// 告警表达式。
	AlarmRuleExpression *string `json:"alarm_rule_expression,omitempty"`

	// 开始报警时间。
	AlarmFirstTime *string `json:"alarm_first_time,omitempty"`

	// 最后一次报警时间。
	AlarmLastTime *string `json:"alarm_last_time,omitempty"`

	// 告警级别。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 唯一告警标识符。
	RestrainKey *string `json:"restrain_key,omitempty"`

	// 告警状态。
	Status *string `json:"status,omitempty"`
}

func (o AlarmDataVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDataVo struct{}"
	}

	return strings.Join([]string{"AlarmDataVo", string(data)}, " ")
}
