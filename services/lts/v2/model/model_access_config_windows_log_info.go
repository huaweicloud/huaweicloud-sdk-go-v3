package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 日志接入采集Windows事件日志
type AccessConfigWindowsLogInfo struct {
	// 采集Windows事件日志类型。Application：应用系统，System：系统，Security：安全，Setup：启动

	Categorys *[]AccessConfigWindowsLogInfoCategorys `json:"categorys,omitempty"`

	TimeOffset *AccessConfigTimeOffset `json:"time_offset,omitempty"`
	// 事件等级。information：info，warning：告警，error：错误，critical：关键，verbose：冗长

	EventLevel *[]AccessConfigWindowsLogInfoEventLevel `json:"event_level,omitempty"`
}

func (o AccessConfigWindowsLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigWindowsLogInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigWindowsLogInfo", string(data)}, " ")
}

type AccessConfigWindowsLogInfoCategorys struct {
	value string
}

type AccessConfigWindowsLogInfoCategorysEnum struct {
	APPLICATION AccessConfigWindowsLogInfoCategorys
	SYSTEM      AccessConfigWindowsLogInfoCategorys
	SECURITY    AccessConfigWindowsLogInfoCategorys
	SETUP       AccessConfigWindowsLogInfoCategorys
}

func GetAccessConfigWindowsLogInfoCategorysEnum() AccessConfigWindowsLogInfoCategorysEnum {
	return AccessConfigWindowsLogInfoCategorysEnum{
		APPLICATION: AccessConfigWindowsLogInfoCategorys{
			value: "Application",
		},
		SYSTEM: AccessConfigWindowsLogInfoCategorys{
			value: "System",
		},
		SECURITY: AccessConfigWindowsLogInfoCategorys{
			value: "Security",
		},
		SETUP: AccessConfigWindowsLogInfoCategorys{
			value: "Setup",
		},
	}
}

func (c AccessConfigWindowsLogInfoCategorys) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoCategorys) UnmarshalJSON(b []byte) error {
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

type AccessConfigWindowsLogInfoEventLevel struct {
	value string
}

type AccessConfigWindowsLogInfoEventLevelEnum struct {
	INFORMATION AccessConfigWindowsLogInfoEventLevel
	WARNING     AccessConfigWindowsLogInfoEventLevel
	ERROR       AccessConfigWindowsLogInfoEventLevel
	CRITICAL    AccessConfigWindowsLogInfoEventLevel
	VERBOSE     AccessConfigWindowsLogInfoEventLevel
}

func GetAccessConfigWindowsLogInfoEventLevelEnum() AccessConfigWindowsLogInfoEventLevelEnum {
	return AccessConfigWindowsLogInfoEventLevelEnum{
		INFORMATION: AccessConfigWindowsLogInfoEventLevel{
			value: "information",
		},
		WARNING: AccessConfigWindowsLogInfoEventLevel{
			value: "warning",
		},
		ERROR: AccessConfigWindowsLogInfoEventLevel{
			value: "error",
		},
		CRITICAL: AccessConfigWindowsLogInfoEventLevel{
			value: "critical",
		},
		VERBOSE: AccessConfigWindowsLogInfoEventLevel{
			value: "verbose",
		},
	}
}

func (c AccessConfigWindowsLogInfoEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoEventLevel) UnmarshalJSON(b []byte) error {
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
