package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmsRequest Request Object
type ListAlarmsRequest struct {

	// 告警记录ID,以ah开头，后跟22位由字母或数字组成的字符串
	AlarmRecordId *string `json:"alarm_record_id,omitempty"`

	// 告警状态，ok为正常，alarm为告警，invalid为已失效
	AlarmStatus *ListAlarmsRequestAlarmStatus `json:"alarm_status,omitempty"`

	// 告警类型
	AlarmType *ListAlarmsRequestAlarmType `json:"alarm_type,omitempty"`

	// 告警资源ID,多值可以以逗号分割,
	ResourceId *string `json:"resource_id,omitempty"`

	// 告警级别，1为紧急，2为重要，3为次要，4为提示
	AlarmLevel *ListAlarmsRequestAlarmLevel `json:"alarm_level,omitempty"`

	// 产生告警开始时间
	From *string `json:"from,omitempty"`

	// 产生告警结束时间
	To *string `json:"to,omitempty"`

	// 分页游标
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}

type ListAlarmsRequestAlarmStatus struct {
	value string
}

type ListAlarmsRequestAlarmStatusEnum struct {
	OK      ListAlarmsRequestAlarmStatus
	ALRAM   ListAlarmsRequestAlarmStatus
	INVALID ListAlarmsRequestAlarmStatus
}

func GetListAlarmsRequestAlarmStatusEnum() ListAlarmsRequestAlarmStatusEnum {
	return ListAlarmsRequestAlarmStatusEnum{
		OK: ListAlarmsRequestAlarmStatus{
			value: "ok",
		},
		ALRAM: ListAlarmsRequestAlarmStatus{
			value: "alram",
		},
		INVALID: ListAlarmsRequestAlarmStatus{
			value: "invalid",
		},
	}
}

func (c ListAlarmsRequestAlarmStatus) Value() string {
	return c.value
}

func (c ListAlarmsRequestAlarmStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRequestAlarmStatus) UnmarshalJSON(b []byte) error {
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

type ListAlarmsRequestAlarmType struct {
	value string
}

type ListAlarmsRequestAlarmTypeEnum struct {
	EVENT  ListAlarmsRequestAlarmType
	METRIC ListAlarmsRequestAlarmType
}

func GetListAlarmsRequestAlarmTypeEnum() ListAlarmsRequestAlarmTypeEnum {
	return ListAlarmsRequestAlarmTypeEnum{
		EVENT: ListAlarmsRequestAlarmType{
			value: "event",
		},
		METRIC: ListAlarmsRequestAlarmType{
			value: "metric",
		},
	}
}

func (c ListAlarmsRequestAlarmType) Value() string {
	return c.value
}

func (c ListAlarmsRequestAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRequestAlarmType) UnmarshalJSON(b []byte) error {
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

type ListAlarmsRequestAlarmLevel struct {
	value int32
}

type ListAlarmsRequestAlarmLevelEnum struct {
	E_1 ListAlarmsRequestAlarmLevel
	E_2 ListAlarmsRequestAlarmLevel
	E_3 ListAlarmsRequestAlarmLevel
	E_4 ListAlarmsRequestAlarmLevel
}

func GetListAlarmsRequestAlarmLevelEnum() ListAlarmsRequestAlarmLevelEnum {
	return ListAlarmsRequestAlarmLevelEnum{
		E_1: ListAlarmsRequestAlarmLevel{
			value: 1,
		}, E_2: ListAlarmsRequestAlarmLevel{
			value: 2,
		}, E_3: ListAlarmsRequestAlarmLevel{
			value: 3,
		}, E_4: ListAlarmsRequestAlarmLevel{
			value: 4,
		},
	}
}

func (c ListAlarmsRequestAlarmLevel) Value() int32 {
	return c.value
}

func (c ListAlarmsRequestAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRequestAlarmLevel) UnmarshalJSON(b []byte) error {
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
