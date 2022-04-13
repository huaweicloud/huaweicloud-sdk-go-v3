package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 日志接入采集Windows事件日志
type AccessConfigWindowsLogInfoCreate struct {
	// 采集Windows事件日志类型。Application：应用系统，System：系统，Security：安全，Setup：启动

	Categorys []AccessConfigWindowsLogInfoCreateCategorys `json:"categorys"`

	TimeOffset *AccessConfigTimeOffsetCreate `json:"time_offset"`
	// 事件等级。information：info，warning：告警，error：错误，critical：关键，verbose：冗长

	EventLevel []AccessConfigWindowsLogInfoCreateEventLevel `json:"event_level"`
}

func (o AccessConfigWindowsLogInfoCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigWindowsLogInfoCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigWindowsLogInfoCreate", string(data)}, " ")
}

type AccessConfigWindowsLogInfoCreateCategorys struct {
	value string
}

type AccessConfigWindowsLogInfoCreateCategorysEnum struct {
	APPLICATION AccessConfigWindowsLogInfoCreateCategorys
	SYSTEM      AccessConfigWindowsLogInfoCreateCategorys
	SECURITY    AccessConfigWindowsLogInfoCreateCategorys
	SETUP       AccessConfigWindowsLogInfoCreateCategorys
}

func GetAccessConfigWindowsLogInfoCreateCategorysEnum() AccessConfigWindowsLogInfoCreateCategorysEnum {
	return AccessConfigWindowsLogInfoCreateCategorysEnum{
		APPLICATION: AccessConfigWindowsLogInfoCreateCategorys{
			value: "Application",
		},
		SYSTEM: AccessConfigWindowsLogInfoCreateCategorys{
			value: "System",
		},
		SECURITY: AccessConfigWindowsLogInfoCreateCategorys{
			value: "Security",
		},
		SETUP: AccessConfigWindowsLogInfoCreateCategorys{
			value: "Setup",
		},
	}
}

func (c AccessConfigWindowsLogInfoCreateCategorys) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoCreateCategorys) UnmarshalJSON(b []byte) error {
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

type AccessConfigWindowsLogInfoCreateEventLevel struct {
	value string
}

type AccessConfigWindowsLogInfoCreateEventLevelEnum struct {
	INFORMATION AccessConfigWindowsLogInfoCreateEventLevel
	WARNING     AccessConfigWindowsLogInfoCreateEventLevel
	ERROR       AccessConfigWindowsLogInfoCreateEventLevel
	CRITICAL    AccessConfigWindowsLogInfoCreateEventLevel
	VERBOSE     AccessConfigWindowsLogInfoCreateEventLevel
}

func GetAccessConfigWindowsLogInfoCreateEventLevelEnum() AccessConfigWindowsLogInfoCreateEventLevelEnum {
	return AccessConfigWindowsLogInfoCreateEventLevelEnum{
		INFORMATION: AccessConfigWindowsLogInfoCreateEventLevel{
			value: "information",
		},
		WARNING: AccessConfigWindowsLogInfoCreateEventLevel{
			value: "warning",
		},
		ERROR: AccessConfigWindowsLogInfoCreateEventLevel{
			value: "error",
		},
		CRITICAL: AccessConfigWindowsLogInfoCreateEventLevel{
			value: "critical",
		},
		VERBOSE: AccessConfigWindowsLogInfoCreateEventLevel{
			value: "verbose",
		},
	}
}

func (c AccessConfigWindowsLogInfoCreateEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoCreateEventLevel) UnmarshalJSON(b []byte) error {
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
