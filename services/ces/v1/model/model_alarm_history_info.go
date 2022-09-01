package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警历史详细信息
type AlarmHistoryInfo struct {

	// 告警规则的ID，如：al1603131199286dzxpqK3Ez。
	AlarmId *string `json:"alarm_id,omitempty" xml:"alarm_id"`

	// 告警规则的名称，如：alarm-test01。
	AlarmName *string `json:"alarm_name,omitempty" xml:"alarm_name"`

	// 告警规则的描述。
	AlarmDescription *string `json:"alarm_description,omitempty" xml:"alarm_description"`

	Metric *MetricInfo `json:"metric,omitempty" xml:"metric"`

	Condition *Condition `json:"condition,omitempty" xml:"condition"`

	// 告警历史的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	AlarmLevel *int32 `json:"alarm_level,omitempty" xml:"alarm_level"`

	// 告警类型； 仅针对事件告警的参数，枚举类型：值为EVENT.SYS或者EVENT.CUSTOM
	AlarmType *string `json:"alarm_type,omitempty" xml:"alarm_type"`

	// 告警规则是否被启用，值为true或者false。
	AlarmEnabled *bool `json:"alarm_enabled,omitempty" xml:"alarm_enabled"`

	// 告警规则的告警触发动作是否被启用，值为true或者false。
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty" xml:"alarm_action_enabled"`

	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notificationList：告警状态发生变化时，被通知对象的列表。
	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty" xml:"alarm_actions"`

	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：  notification：通知。  notificationList：告警状态发生变化时，被通知对象的列表。
	OkActions *[]AlarmActions `json:"ok_actions,omitempty" xml:"ok_actions"`

	// 数据不足触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： 数据不足触发告警通知类型，取值为notification。 notificationList：数据不足触发告警通知时，被通知对象的ID列表。
	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty" xml:"insufficientdata_actions"`

	// 告警状态变更的时间，UNIX时间戳，单位毫秒，如：1603131199000。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 企业项目ID。 值为all_granted_eps时，表示所有企业项目；值为0时，表示默认的企业项目default。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 计算出该条告警历史的时间，UNIX时间戳，单位毫秒，如：1603131199469。
	TriggerTime *int64 `json:"trigger_time,omitempty" xml:"trigger_time"`

	// 告警历史的状态，取值为ok，alarm，insufficient_data； ok为正常，alarm为告警，insufficient_data为数据不足。
	AlarmStatus *string `json:"alarm_status,omitempty" xml:"alarm_status"`

	// 计算出该条告警历史的资源监控数据的一组数据上报时间和监控数值。
	Datapoints *[]DataPointForAlarmHistory `json:"datapoints,omitempty" xml:"datapoints"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty" xml:"additional_info"`
}

func (o AlarmHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryInfo struct{}"
	}

	return strings.Join([]string{"AlarmHistoryInfo", string(data)}, " ")
}
