package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MetricAlarms struct {
	// 告警名称。

	AlarmName string `json:"alarm_name"`
	// 告警描述。

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *MetricInfoForAlarm `json:"metric"`

	Condition *Condition `json:"condition"`
	// 是否启用该条告警。

	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`
	// 告警级别，默认为2，级别为1、2、3、4。分别对应紧急、重要、次要、提示。

	AlarmLevel *int32 `json:"alarm_level,omitempty"`
	// 告警类型。 仅针对事件告警的参数，枚举类型：EVENT.SYS或者EVENT.CUSTOM

	AlarmType *MetricAlarmsAlarmType `json:"alarm_type,omitempty"`
	// 是否启用该条告警触发的动作。

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notificationList：告警状态发生变化时，被通知对象的列表。

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`
	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：  notification：通知。  notificationList：告警状态发生变化时，被通知对象的列表。

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`
	// 数据不足触发的动作。  结构如下：  {  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： 数据不足触发告警通知类型，取值为notification。 notificationList：数据不足触发告警通知时，被通知对象的ID列表。

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`
	// 告警规则生效的开始时间，告警规则仅在生效时间内发送通知消息。例如alarm_action_begin_time为8:00，alarm_action_end_time为20:00时，则对应的告警规则仅在08:00-20:00发送通知消息。

	AlarmActionBeginTime *string `json:"alarm_action_begin_time,omitempty"`
	// 告警规则生效的结束时间，告警规则仅在生效时间内发送通知消息。例如alarm_action_begin_time为8:00，alarm_action_end_time为20:00时，则对应的告警规则仅在08:00-20:00发送通知消息。

	AlarmActionEndTime *string `json:"alarm_action_end_time,omitempty"`
	// 告警规则的ID。

	AlarmId string `json:"alarm_id"`
	// 告警状态变更的时间，UNIX时间戳，单位毫秒。

	UpdateTime int64 `json:"update_time"`
	// 告警状态，取值说明：  ok，正常 alarm，告警 insufficient_data，数据不足

	AlarmState string `json:"alarm_state"`
	// 企业项目ID。 取值为all_granted_eps时，表示所有企业项目; 取值为0时，表示默认的企业项目default。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o MetricAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricAlarms struct{}"
	}

	return strings.Join([]string{"MetricAlarms", string(data)}, " ")
}

type MetricAlarmsAlarmType struct {
	value string
}

type MetricAlarmsAlarmTypeEnum struct {
	EVENT_SYS    MetricAlarmsAlarmType
	EVENT_CUSTOM MetricAlarmsAlarmType
}

func GetMetricAlarmsAlarmTypeEnum() MetricAlarmsAlarmTypeEnum {
	return MetricAlarmsAlarmTypeEnum{
		EVENT_SYS: MetricAlarmsAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: MetricAlarmsAlarmType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c MetricAlarmsAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricAlarmsAlarmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
