package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmConfigResponse Response Object
type ListAlarmConfigResponse struct {

	// 告警类型
	AlarmType *ListAlarmConfigResponseAlarmType `json:"alarm_type,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 是否发送短信
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 是否开启每日告警
	DailyAlarm *bool `json:"daily_alarm,omitempty"`

	// 是否开启实时告警
	RealtimeAlarm *bool `json:"realtime_alarm,omitempty"`

	// 告警等级
	AlarmLevel *[]ListAlarmConfigResponseAlarmLevel `json:"alarm_level,omitempty"`

	// 忽略的事件列表
	IgnoredEventClassList *[]ListAlarmConfigResponseIgnoredEventClassList `json:"ignored_event_class_list,omitempty"`
	HttpStatusCode        int                                             `json:"-"`
}

func (o ListAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmConfigResponse", string(data)}, " ")
}

type ListAlarmConfigResponseAlarmType struct {
	value string
}

type ListAlarmConfigResponseAlarmTypeEnum struct {
	SMS ListAlarmConfigResponseAlarmType
	RSS ListAlarmConfigResponseAlarmType
}

func GetListAlarmConfigResponseAlarmTypeEnum() ListAlarmConfigResponseAlarmTypeEnum {
	return ListAlarmConfigResponseAlarmTypeEnum{
		SMS: ListAlarmConfigResponseAlarmType{
			value: "sms",
		},
		RSS: ListAlarmConfigResponseAlarmType{
			value: "rss",
		},
	}
}

func (c ListAlarmConfigResponseAlarmType) Value() string {
	return c.value
}

func (c ListAlarmConfigResponseAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmConfigResponseAlarmType) UnmarshalJSON(b []byte) error {
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

type ListAlarmConfigResponseAlarmLevel struct {
	value string
}

type ListAlarmConfigResponseAlarmLevelEnum struct {
	INFO     ListAlarmConfigResponseAlarmLevel
	LOW      ListAlarmConfigResponseAlarmLevel
	MEDIUM   ListAlarmConfigResponseAlarmLevel
	HIGH     ListAlarmConfigResponseAlarmLevel
	CRITICAL ListAlarmConfigResponseAlarmLevel
}

func GetListAlarmConfigResponseAlarmLevelEnum() ListAlarmConfigResponseAlarmLevelEnum {
	return ListAlarmConfigResponseAlarmLevelEnum{
		INFO: ListAlarmConfigResponseAlarmLevel{
			value: "info",
		},
		LOW: ListAlarmConfigResponseAlarmLevel{
			value: "low",
		},
		MEDIUM: ListAlarmConfigResponseAlarmLevel{
			value: "medium",
		},
		HIGH: ListAlarmConfigResponseAlarmLevel{
			value: "high",
		},
		CRITICAL: ListAlarmConfigResponseAlarmLevel{
			value: "critical",
		},
	}
}

func (c ListAlarmConfigResponseAlarmLevel) Value() string {
	return c.value
}

func (c ListAlarmConfigResponseAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmConfigResponseAlarmLevel) UnmarshalJSON(b []byte) error {
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

type ListAlarmConfigResponseIgnoredEventClassList struct {
	value string
}

type ListAlarmConfigResponseIgnoredEventClassListEnum struct {
	DANGEROUS_PORTS          ListAlarmConfigResponseIgnoredEventClassList
	CRITICAL_VULNERABILITIES ListAlarmConfigResponseIgnoredEventClassList
	WEAK_PASSWORDS           ListAlarmConfigResponseIgnoredEventClassList
	UNSAFE_CONFIGURATIONS    ListAlarmConfigResponseIgnoredEventClassList
	RASP                     ListAlarmConfigResponseIgnoredEventClassList
	SECRASP                  ListAlarmConfigResponseIgnoredEventClassList
	E_8002                   ListAlarmConfigResponseIgnoredEventClassList
	E_8003                   ListAlarmConfigResponseIgnoredEventClassList
	E_1001                   ListAlarmConfigResponseIgnoredEventClassList
	E_1010                   ListAlarmConfigResponseIgnoredEventClassList
	E_1011                   ListAlarmConfigResponseIgnoredEventClassList
	E_1015                   ListAlarmConfigResponseIgnoredEventClassList
	E_1017                   ListAlarmConfigResponseIgnoredEventClassList
	E_2047                   ListAlarmConfigResponseIgnoredEventClassList
	E_2048                   ListAlarmConfigResponseIgnoredEventClassList
	E_2049                   ListAlarmConfigResponseIgnoredEventClassList
	E_3002                   ListAlarmConfigResponseIgnoredEventClassList
	E_3003                   ListAlarmConfigResponseIgnoredEventClassList
	E_3004                   ListAlarmConfigResponseIgnoredEventClassList
	E_3005                   ListAlarmConfigResponseIgnoredEventClassList
	E_3007                   ListAlarmConfigResponseIgnoredEventClassList
	E_3015                   ListAlarmConfigResponseIgnoredEventClassList
	E_3018                   ListAlarmConfigResponseIgnoredEventClassList
	E_3027                   ListAlarmConfigResponseIgnoredEventClassList
	E_4002                   ListAlarmConfigResponseIgnoredEventClassList
	E_4007                   ListAlarmConfigResponseIgnoredEventClassList
	E_4004                   ListAlarmConfigResponseIgnoredEventClassList
	E_4006                   ListAlarmConfigResponseIgnoredEventClassList
	E_2042                   ListAlarmConfigResponseIgnoredEventClassList
	E_2044                   ListAlarmConfigResponseIgnoredEventClassList
	E_3008                   ListAlarmConfigResponseIgnoredEventClassList
	E_3009                   ListAlarmConfigResponseIgnoredEventClassList
	E_3016                   ListAlarmConfigResponseIgnoredEventClassList
	E_3021                   ListAlarmConfigResponseIgnoredEventClassList
}

func GetListAlarmConfigResponseIgnoredEventClassListEnum() ListAlarmConfigResponseIgnoredEventClassListEnum {
	return ListAlarmConfigResponseIgnoredEventClassListEnum{
		DANGEROUS_PORTS: ListAlarmConfigResponseIgnoredEventClassList{
			value: "dangerous_ports",
		},
		CRITICAL_VULNERABILITIES: ListAlarmConfigResponseIgnoredEventClassList{
			value: "critical_vulnerabilities",
		},
		WEAK_PASSWORDS: ListAlarmConfigResponseIgnoredEventClassList{
			value: "weak_passwords",
		},
		UNSAFE_CONFIGURATIONS: ListAlarmConfigResponseIgnoredEventClassList{
			value: "unsafe_configurations",
		},
		RASP: ListAlarmConfigResponseIgnoredEventClassList{
			value: "RASP",
		},
		SECRASP: ListAlarmConfigResponseIgnoredEventClassList{
			value: "SECRASP",
		},
		E_8002: ListAlarmConfigResponseIgnoredEventClassList{
			value: "8002",
		},
		E_8003: ListAlarmConfigResponseIgnoredEventClassList{
			value: "8003",
		},
		E_1001: ListAlarmConfigResponseIgnoredEventClassList{
			value: "1001",
		},
		E_1010: ListAlarmConfigResponseIgnoredEventClassList{
			value: "1010",
		},
		E_1011: ListAlarmConfigResponseIgnoredEventClassList{
			value: "1011",
		},
		E_1015: ListAlarmConfigResponseIgnoredEventClassList{
			value: "1015",
		},
		E_1017: ListAlarmConfigResponseIgnoredEventClassList{
			value: "1017",
		},
		E_2047: ListAlarmConfigResponseIgnoredEventClassList{
			value: "2047",
		},
		E_2048: ListAlarmConfigResponseIgnoredEventClassList{
			value: "2048",
		},
		E_2049: ListAlarmConfigResponseIgnoredEventClassList{
			value: "2049",
		},
		E_3002: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3002",
		},
		E_3003: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3003",
		},
		E_3004: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3004",
		},
		E_3005: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3005",
		},
		E_3007: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3007",
		},
		E_3015: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3015",
		},
		E_3018: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3018",
		},
		E_3027: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3027",
		},
		E_4002: ListAlarmConfigResponseIgnoredEventClassList{
			value: "4002",
		},
		E_4007: ListAlarmConfigResponseIgnoredEventClassList{
			value: "4007",
		},
		E_4004: ListAlarmConfigResponseIgnoredEventClassList{
			value: "4004",
		},
		E_4006: ListAlarmConfigResponseIgnoredEventClassList{
			value: "4006",
		},
		E_2042: ListAlarmConfigResponseIgnoredEventClassList{
			value: "2042",
		},
		E_2044: ListAlarmConfigResponseIgnoredEventClassList{
			value: "2044",
		},
		E_3008: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3008",
		},
		E_3009: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3009",
		},
		E_3016: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3016",
		},
		E_3021: ListAlarmConfigResponseIgnoredEventClassList{
			value: "3021",
		},
	}
}

func (c ListAlarmConfigResponseIgnoredEventClassList) Value() string {
	return c.value
}

func (c ListAlarmConfigResponseIgnoredEventClassList) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmConfigResponseIgnoredEventClassList) UnmarshalJSON(b []byte) error {
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
