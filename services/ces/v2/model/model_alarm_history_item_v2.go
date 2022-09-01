package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 告警记录详细信息
type AlarmHistoryItemV2 struct {

	// 告警记录ID
	RecordId *string `json:"record_id,omitempty" xml:"record_id"`

	// 告警规则的ID，如：al1603131199286dzxpqK3Ez。
	AlarmId *string `json:"alarm_id,omitempty" xml:"alarm_id"`

	// 告警规则的名称，如：alarm-test01。
	Name *string `json:"name,omitempty" xml:"name"`

	// 告警记录的状态，取值为ok，alarm，invalid； ok为正常，alarm为告警，invalid为已失效。
	Status *AlarmHistoryItemV2Status `json:"status,omitempty" xml:"status"`

	// 告警记录的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	Level *AlarmHistoryItemV2Level `json:"level,omitempty" xml:"level"`

	// 告警类型； 仅针对事件告警的参数，枚举类型：值为EVENT.SYS或者EVENT.CUSTOM
	Type *AlarmHistoryItemV2Type `json:"type,omitempty" xml:"type"`

	// 是否发送通知，值为true或者false。
	ActionEnabled *bool `json:"action_enabled,omitempty" xml:"action_enabled"`

	// 产生时间,UTC时间
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty" xml:"begin_time"`

	// 结束时间，UTC时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty" xml:"end_time"`

	Metric *Metric `json:"metric,omitempty" xml:"metric"`

	Condition *AlarmCondition `json:"condition,omitempty" xml:"condition"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty" xml:"additional_info"`

	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notification_list\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notification_list：告警状态发生变化时，被通知对象的列表。
	AlarmActions *[]Notification `json:"alarm_actions,omitempty" xml:"alarm_actions"`

	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notification_list\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：  notification：通知。  notification_list：告警状态发生变化时，被通知对象的列表。
	OkActions *[]Notification `json:"ok_actions,omitempty" xml:"ok_actions"`

	// 计算出该条告警记录的资源监控数据上报时间和监控数值。
	Datapoints *[]DataPointInfo `json:"datapoints,omitempty" xml:"datapoints"`
}

func (o AlarmHistoryItemV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2 struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2", string(data)}, " ")
}

type AlarmHistoryItemV2Status struct {
	value string
}

type AlarmHistoryItemV2StatusEnum struct {
	OK      AlarmHistoryItemV2Status
	ALARM   AlarmHistoryItemV2Status
	INVALID AlarmHistoryItemV2Status
}

func GetAlarmHistoryItemV2StatusEnum() AlarmHistoryItemV2StatusEnum {
	return AlarmHistoryItemV2StatusEnum{
		OK: AlarmHistoryItemV2Status{
			value: "ok",
		},
		ALARM: AlarmHistoryItemV2Status{
			value: "alarm",
		},
		INVALID: AlarmHistoryItemV2Status{
			value: "invalid",
		},
	}
}

func (c AlarmHistoryItemV2Status) Value() string {
	return c.value
}

func (c AlarmHistoryItemV2Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Status) UnmarshalJSON(b []byte) error {
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

type AlarmHistoryItemV2Level struct {
	value int32
}

type AlarmHistoryItemV2LevelEnum struct {
	E_1 AlarmHistoryItemV2Level
	E_2 AlarmHistoryItemV2Level
	E_3 AlarmHistoryItemV2Level
	E_4 AlarmHistoryItemV2Level
}

func GetAlarmHistoryItemV2LevelEnum() AlarmHistoryItemV2LevelEnum {
	return AlarmHistoryItemV2LevelEnum{
		E_1: AlarmHistoryItemV2Level{
			value: 1,
		}, E_2: AlarmHistoryItemV2Level{
			value: 2,
		}, E_3: AlarmHistoryItemV2Level{
			value: 3,
		}, E_4: AlarmHistoryItemV2Level{
			value: 4,
		},
	}
}

func (c AlarmHistoryItemV2Level) Value() int32 {
	return c.value
}

func (c AlarmHistoryItemV2Level) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Level) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AlarmHistoryItemV2Type struct {
	value string
}

type AlarmHistoryItemV2TypeEnum struct {
	EVENT_SYS    AlarmHistoryItemV2Type
	EVENT_CUSTOM AlarmHistoryItemV2Type
}

func GetAlarmHistoryItemV2TypeEnum() AlarmHistoryItemV2TypeEnum {
	return AlarmHistoryItemV2TypeEnum{
		EVENT_SYS: AlarmHistoryItemV2Type{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: AlarmHistoryItemV2Type{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c AlarmHistoryItemV2Type) Value() string {
	return c.value
}

func (c AlarmHistoryItemV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Type) UnmarshalJSON(b []byte) error {
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
