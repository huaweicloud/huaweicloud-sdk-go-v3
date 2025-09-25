package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAlarmConfigRequestInfo alarm config
type UpdateAlarmConfigRequestInfo struct {

	// 告警类型
	AlarmType *UpdateAlarmConfigRequestInfoAlarmType `json:"alarm_type,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 是否发送短信
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 是否开启每日告警
	DailyAlarm *bool `json:"daily_alarm,omitempty"`

	// 是否开启实时告警
	RealtimeAlarm *bool `json:"realtime_alarm,omitempty"`

	// 告警等级
	AlarmLevel *UpdateAlarmConfigRequestInfoAlarmLevel `json:"alarm_level,omitempty"`

	// 忽略事件列表
	IgnoredEventClassList *[]UpdateAlarmConfigRequestInfoIgnoredEventClassList `json:"ignored_event_class_list,omitempty"`
}

func (o UpdateAlarmConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateAlarmConfigRequestInfo", string(data)}, " ")
}

type UpdateAlarmConfigRequestInfoAlarmType struct {
	value string
}

type UpdateAlarmConfigRequestInfoAlarmTypeEnum struct {
	SMS UpdateAlarmConfigRequestInfoAlarmType
	RSS UpdateAlarmConfigRequestInfoAlarmType
}

func GetUpdateAlarmConfigRequestInfoAlarmTypeEnum() UpdateAlarmConfigRequestInfoAlarmTypeEnum {
	return UpdateAlarmConfigRequestInfoAlarmTypeEnum{
		SMS: UpdateAlarmConfigRequestInfoAlarmType{
			value: "sms",
		},
		RSS: UpdateAlarmConfigRequestInfoAlarmType{
			value: "rss",
		},
	}
}

func (c UpdateAlarmConfigRequestInfoAlarmType) Value() string {
	return c.value
}

func (c UpdateAlarmConfigRequestInfoAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmConfigRequestInfoAlarmType) UnmarshalJSON(b []byte) error {
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

type UpdateAlarmConfigRequestInfoAlarmLevel struct {
	value string
}

type UpdateAlarmConfigRequestInfoAlarmLevelEnum struct {
	INFO     UpdateAlarmConfigRequestInfoAlarmLevel
	LOW      UpdateAlarmConfigRequestInfoAlarmLevel
	MEDIUM   UpdateAlarmConfigRequestInfoAlarmLevel
	HIGH     UpdateAlarmConfigRequestInfoAlarmLevel
	CRITICAL UpdateAlarmConfigRequestInfoAlarmLevel
}

func GetUpdateAlarmConfigRequestInfoAlarmLevelEnum() UpdateAlarmConfigRequestInfoAlarmLevelEnum {
	return UpdateAlarmConfigRequestInfoAlarmLevelEnum{
		INFO: UpdateAlarmConfigRequestInfoAlarmLevel{
			value: "info",
		},
		LOW: UpdateAlarmConfigRequestInfoAlarmLevel{
			value: "low",
		},
		MEDIUM: UpdateAlarmConfigRequestInfoAlarmLevel{
			value: "medium",
		},
		HIGH: UpdateAlarmConfigRequestInfoAlarmLevel{
			value: "high",
		},
		CRITICAL: UpdateAlarmConfigRequestInfoAlarmLevel{
			value: "critical",
		},
	}
}

func (c UpdateAlarmConfigRequestInfoAlarmLevel) Value() string {
	return c.value
}

func (c UpdateAlarmConfigRequestInfoAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmConfigRequestInfoAlarmLevel) UnmarshalJSON(b []byte) error {
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

type UpdateAlarmConfigRequestInfoIgnoredEventClassList struct {
	value string
}

type UpdateAlarmConfigRequestInfoIgnoredEventClassListEnum struct {
	DANGEROUS_PORTS          UpdateAlarmConfigRequestInfoIgnoredEventClassList
	CRITICAL_VULNERABILITIES UpdateAlarmConfigRequestInfoIgnoredEventClassList
	WEAK_PASSWORDS           UpdateAlarmConfigRequestInfoIgnoredEventClassList
	UNSAFE_CONFIGURATIONS    UpdateAlarmConfigRequestInfoIgnoredEventClassList
	RASP                     UpdateAlarmConfigRequestInfoIgnoredEventClassList
	SECRASP                  UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_8002                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_8003                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_1001                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_1010                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_1011                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_1015                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_1017                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_2047                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_2048                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_2049                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3002                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3003                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3004                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3005                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3007                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3015                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3018                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3027                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_4002                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_4007                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_4004                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_4006                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_2042                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_2044                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3008                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3009                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3016                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
	E_3021                   UpdateAlarmConfigRequestInfoIgnoredEventClassList
}

func GetUpdateAlarmConfigRequestInfoIgnoredEventClassListEnum() UpdateAlarmConfigRequestInfoIgnoredEventClassListEnum {
	return UpdateAlarmConfigRequestInfoIgnoredEventClassListEnum{
		DANGEROUS_PORTS: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "dangerous_ports",
		},
		CRITICAL_VULNERABILITIES: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "critical_vulnerabilities",
		},
		WEAK_PASSWORDS: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "weak_passwords",
		},
		UNSAFE_CONFIGURATIONS: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "unsafe_configurations",
		},
		RASP: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "RASP",
		},
		SECRASP: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "SECRASP",
		},
		E_8002: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "8002",
		},
		E_8003: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "8003",
		},
		E_1001: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "1001",
		},
		E_1010: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "1010",
		},
		E_1011: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "1011",
		},
		E_1015: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "1015",
		},
		E_1017: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "1017",
		},
		E_2047: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "2047",
		},
		E_2048: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "2048",
		},
		E_2049: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "2049",
		},
		E_3002: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3002",
		},
		E_3003: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3003",
		},
		E_3004: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3004",
		},
		E_3005: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3005",
		},
		E_3007: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3007",
		},
		E_3015: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3015",
		},
		E_3018: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3018",
		},
		E_3027: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3027",
		},
		E_4002: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "4002",
		},
		E_4007: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "4007",
		},
		E_4004: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "4004",
		},
		E_4006: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "4006",
		},
		E_2042: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "2042",
		},
		E_2044: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "2044",
		},
		E_3008: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3008",
		},
		E_3009: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3009",
		},
		E_3016: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3016",
		},
		E_3021: UpdateAlarmConfigRequestInfoIgnoredEventClassList{
			value: "3021",
		},
	}
}

func (c UpdateAlarmConfigRequestInfoIgnoredEventClassList) Value() string {
	return c.value
}

func (c UpdateAlarmConfigRequestInfoIgnoredEventClassList) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmConfigRequestInfoIgnoredEventClassList) UnmarshalJSON(b []byte) error {
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
