package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExpressConfig struct {

	// 同步工作流执行时记录LTS的日志级别，NONE、ERROR\\ALL，默认NONE
	LogLevel *ExpressConfigLogLevel `json:"log_level,omitempty"`

	// 同步工作流是否支持匿名访问
	SupportAnonymous *bool `json:"support_anonymous,omitempty"`
}

func (o ExpressConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpressConfig struct{}"
	}

	return strings.Join([]string{"ExpressConfig", string(data)}, " ")
}

type ExpressConfigLogLevel struct {
	value string
}

type ExpressConfigLogLevelEnum struct {
	NONE  ExpressConfigLogLevel
	ERROR ExpressConfigLogLevel
	ALL   ExpressConfigLogLevel
}

func GetExpressConfigLogLevelEnum() ExpressConfigLogLevelEnum {
	return ExpressConfigLogLevelEnum{
		NONE: ExpressConfigLogLevel{
			value: "NONE",
		},
		ERROR: ExpressConfigLogLevel{
			value: "ERROR",
		},
		ALL: ExpressConfigLogLevel{
			value: "ALL",
		},
	}
}

func (c ExpressConfigLogLevel) Value() string {
	return c.value
}

func (c ExpressConfigLogLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExpressConfigLogLevel) UnmarshalJSON(b []byte) error {
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
