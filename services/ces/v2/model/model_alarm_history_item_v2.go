package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AlarmHistoryItemV2 告警记录详细信息
type AlarmHistoryItemV2 struct {

	// 告警记录ID
	RecordId *string `json:"record_id,omitempty"`

	// 告警规则的ID，如：al1603131199286dzxpqK3Ez。
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警规则的名称，如：alarm-test01。
	Name *string `json:"name,omitempty"`

	// 告警记录的状态，取值为ok，alarm，invalid； ok为正常，alarm为告警，invalid为已失效。
	Status *AlarmHistoryItemV2Status `json:"status,omitempty"`

	// 告警记录的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	Level *AlarmHistoryItemV2Level `json:"level,omitempty"`

	Type *AlarmType `json:"type,omitempty"`

	// 是否发送通知，值为true或者false。
	ActionEnabled *bool `json:"action_enabled,omitempty"`

	// 产生时间,UTC时间
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 结束时间，UTC时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Metric *Metric `json:"metric,omitempty"`

	Condition *AlarmCondition `json:"condition,omitempty"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`

	// 告警触发的动作。  结构如下：  {  \"type\": \"notification\", \"notification_list\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值： notification：通知。 autoscaling：弹性伸缩。 notification_list：告警状态发生变化时，被通知对象的列表。
	AlarmActions *[]Notification `json:"alarm_actions,omitempty"`

	// 告警恢复触发的动作。  结构如下：  {  \"type\": \"notification\", \"notification_list\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：  notification：通知。  notification_list：告警状态发生变化时，被通知对象的列表。
	OkActions *[]Notification `json:"ok_actions,omitempty"`

	// 计算出该条告警记录的资源监控数据上报时间和监控数值。
	DataPoints *[]DataPointInfo `json:"data_points,omitempty"`
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
