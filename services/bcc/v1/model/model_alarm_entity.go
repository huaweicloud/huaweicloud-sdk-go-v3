package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmEntity 告警实体类
type AlarmEntity struct {

	// 告警记录ID
	AlarmRecordId string `json:"alarm_record_id"`

	// 企业项目ID
	EpsId *string `json:"eps_id,omitempty"`

	Status *AlarmStatusEnum `json:"status,omitempty"`

	// 告警级别,取值范围：1,2,3,4
	Severity *string `json:"severity,omitempty"`

	// 告警产生时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 最后更新时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 持续时长
	Duration *string `json:"duration,omitempty"`

	AlarmType *AlarmTypeEnum `json:"alarm_type,omitempty"`

	Namespace *AlarmNamespaceEnum `json:"namespace,omitempty"`

	// 异常资源数量
	AbnormalResource *string `json:"abnormal_resource,omitempty"`

	// 告警策略
	AlarmPolicy *string `json:"alarm_policy,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 告警规则名称
	AlarmRuleName *string `json:"alarm_rule_name,omitempty"`

	// 告警规则ID
	AlarmRuleId string `json:"alarm_rule_id"`

	// 计算出该条告警记录的资源监控数据上报时间和监控数值
	AlarmDatapoints *string `json:"alarm_datapoints,omitempty"`

	// 告警指标
	Metric *string `json:"metric,omitempty"`

	// 告警触发条件
	Condition *string `json:"condition,omitempty"`

	// 告警触发的动作
	AlarmActions *string `json:"alarm_actions,omitempty"`

	// 告警恢复触发的动作
	OkActions *string `json:"ok_actions,omitempty"`

	// 告警记录额外字段
	AdditionalInfo *string `json:"additional_info,omitempty"`
}

func (o AlarmEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEntity struct{}"
	}

	return strings.Join([]string{"AlarmEntity", string(data)}, " ")
}
