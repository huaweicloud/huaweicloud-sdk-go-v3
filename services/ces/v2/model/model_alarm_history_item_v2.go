package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警历史详细信息
type AlarmHistoryItemV2 struct {
	// 告警历史ID

	RecordId *string `json:"record_id,omitempty"`
	// 告警规则的ID，如：al1603131199286dzxpqK3Ez。

	AlarmId *string `json:"alarm_id,omitempty"`
	// 告警规则的名称，如：alarm-test01。

	Name *string `json:"name,omitempty"`
	// 告警历史的状态，取值为ok，alarm，insufficient_data； ok为正常，alarm为告警，insufficient_data数据不足。

	Status *string `json:"status,omitempty"`
	// 告警历史的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。

	Level *int32 `json:"level,omitempty"`
	// 告警类型； 仅针对事件告警的参数，枚举类型：值为EVENT.SYS或者EVENT.CUSTOM

	Type *string `json:"type,omitempty"`
	// 是否发送通知，值为true或者false。

	ActionEnabled *bool `json:"action_enabled,omitempty"`
	// 产生时间,UTC时间

	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`
	// 结束时间，UTC时间

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Metric *Metric `json:"metric,omitempty"`

	Condition *AlarmCondition `json:"condition,omitempty"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notificationList：告警状态发生变化时，被通知对象的列表。

	AlarmActions *[]SmnAction `json:"alarm_actions,omitempty"`
	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：  notification：通知。  notificationList：告警状态发生变化时，被通知对象的列表。

	OkActions *[]SmnAction `json:"ok_actions,omitempty"`
	// 计算出该条告警历史的资源监控数据上报时间和监控数值。

	DataPoints *[]interface{} `json:"data_points,omitempty"`
}

func (o AlarmHistoryItemV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2 struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2", string(data)}, " ")
}
