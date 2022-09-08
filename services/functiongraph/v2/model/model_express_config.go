package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 快速模式配置项
type ExpressConfig struct {

	// 快速模式相关配置，仅在mode配置为EXPRESS时生效 快速模式下流程的执行日志级别，当前支持： ALL: 记录所有节点的执行日志 ERROR：仅记录异常节点执行日志 NONE：不记录日志 注意：当配置为ALL和ERROR级别时租户需要开启LTS相关权限
	LogLevel *ExpressConfigLogLevel `json:"log_level,omitempty"`
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
	ALL   ExpressConfigLogLevel
	ERROR ExpressConfigLogLevel
	NONE  ExpressConfigLogLevel
}

func GetExpressConfigLogLevelEnum() ExpressConfigLogLevelEnum {
	return ExpressConfigLogLevelEnum{
		ALL: ExpressConfigLogLevel{
			value: "ALL",
		},
		ERROR: ExpressConfigLogLevel{
			value: "ERROR",
		},
		NONE: ExpressConfigLogLevel{
			value: "NONE",
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
