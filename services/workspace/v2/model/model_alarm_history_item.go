package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmHistoryItem struct {

	// 告警记录。
	RecordId *string `json:"record_id,omitempty"`

	// 告警规则ID。
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警规则的名称。
	Name *string `json:"name,omitempty"`

	// 告警记录的状态，取值为ok，alarm，invalid； ok为正常，alarm为告警，invalid为已失效。
	Status *string `json:"status,omitempty"`

	// 告警规则类型 | ALL_INSTANCE为全部资源指标告警， RESOURCE_GROUP为资源分组指标告警， MULTI_INSTANCE为指定资源指标告警， EVENT.SYS为系统事件告警， EVENT.CUSTOM自定义事件告警， DNSHealthCheck为健康检查告警。
	Type *string `json:"type,omitempty"`

	// 告警记录的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	Level *int32 `json:"level,omitempty"`

	// 产生时间,UTC时间。
	BeginTime *string `json:"begin_time,omitempty"`

	Metric *AlarmMetric `json:"metric,omitempty"`

	Condition *AlarmCondition `json:"condition,omitempty"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`

	// 计算出该条告警记录的资源监控数据上报时间和监控数值。
	DataPoints *[]DataPointInfo `json:"data_points,omitempty"`
}

func (o AlarmHistoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItem struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItem", string(data)}, " ")
}
