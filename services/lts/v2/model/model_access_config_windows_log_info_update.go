package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigWindowsLogInfoUpdate 日志接入采集Windows事件日志
type AccessConfigWindowsLogInfoUpdate struct {

	// 采集Windows事件日志类型。Application：应用系统，System：系统，Security：安全，Setup：启动
	Categorys *[]AccessConfigWindowsLogInfoUpdateCategorys `json:"categorys,omitempty"`

	TimeOffset *AccessConfigTimeOffsetCreate `json:"time_offset,omitempty"`

	// 事件等级。information：info，warning：告警，error：错误，critical：关键，verbose：冗长
	EventLevel *[]AccessConfigWindowsLogInfoUpdateEventLevel `json:"event_level,omitempty"`
}

func (o AccessConfigWindowsLogInfoUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigWindowsLogInfoUpdate struct{}"
	}

	return strings.Join([]string{"AccessConfigWindowsLogInfoUpdate", string(data)}, " ")
}

type AccessConfigWindowsLogInfoUpdateCategorys struct {
	value string
}

type AccessConfigWindowsLogInfoUpdateCategorysEnum struct {
	APPLICATION AccessConfigWindowsLogInfoUpdateCategorys
	SYSTEM      AccessConfigWindowsLogInfoUpdateCategorys
	SECURITY    AccessConfigWindowsLogInfoUpdateCategorys
	SETUP       AccessConfigWindowsLogInfoUpdateCategorys
}

func GetAccessConfigWindowsLogInfoUpdateCategorysEnum() AccessConfigWindowsLogInfoUpdateCategorysEnum {
	return AccessConfigWindowsLogInfoUpdateCategorysEnum{
		APPLICATION: AccessConfigWindowsLogInfoUpdateCategorys{
			value: "Application",
		},
		SYSTEM: AccessConfigWindowsLogInfoUpdateCategorys{
			value: "System",
		},
		SECURITY: AccessConfigWindowsLogInfoUpdateCategorys{
			value: "Security",
		},
		SETUP: AccessConfigWindowsLogInfoUpdateCategorys{
			value: "Setup",
		},
	}
}

func (c AccessConfigWindowsLogInfoUpdateCategorys) Value() string {
	return c.value
}

func (c AccessConfigWindowsLogInfoUpdateCategorys) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoUpdateCategorys) UnmarshalJSON(b []byte) error {
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

type AccessConfigWindowsLogInfoUpdateEventLevel struct {
	value string
}

type AccessConfigWindowsLogInfoUpdateEventLevelEnum struct {
	INFORMATION AccessConfigWindowsLogInfoUpdateEventLevel
	WARNING     AccessConfigWindowsLogInfoUpdateEventLevel
	ERROR       AccessConfigWindowsLogInfoUpdateEventLevel
	CRITICAL    AccessConfigWindowsLogInfoUpdateEventLevel
	VERBOSE     AccessConfigWindowsLogInfoUpdateEventLevel
}

func GetAccessConfigWindowsLogInfoUpdateEventLevelEnum() AccessConfigWindowsLogInfoUpdateEventLevelEnum {
	return AccessConfigWindowsLogInfoUpdateEventLevelEnum{
		INFORMATION: AccessConfigWindowsLogInfoUpdateEventLevel{
			value: "information",
		},
		WARNING: AccessConfigWindowsLogInfoUpdateEventLevel{
			value: "warning",
		},
		ERROR: AccessConfigWindowsLogInfoUpdateEventLevel{
			value: "error",
		},
		CRITICAL: AccessConfigWindowsLogInfoUpdateEventLevel{
			value: "critical",
		},
		VERBOSE: AccessConfigWindowsLogInfoUpdateEventLevel{
			value: "verbose",
		},
	}
}

func (c AccessConfigWindowsLogInfoUpdateEventLevel) Value() string {
	return c.value
}

func (c AccessConfigWindowsLogInfoUpdateEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigWindowsLogInfoUpdateEventLevel) UnmarshalJSON(b []byte) error {
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
