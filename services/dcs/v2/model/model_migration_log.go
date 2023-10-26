package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MigrationLog 数据迁移进度日志信息
type MigrationLog struct {

	// 迁移日志生成时间，形如：2023-05-15T09:11:25.449Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 日志级别
	LogLevel *MigrationLogLogLevel `json:"log_level,omitempty"`

	// 日志信息
	Message *string `json:"message,omitempty"`

	// 日志的编码
	LogCode *string `json:"log_code,omitempty"`

	// 日志中的关键字
	Keyword *[]string `json:"keyword,omitempty"`
}

func (o MigrationLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationLog struct{}"
	}

	return strings.Join([]string{"MigrationLog", string(data)}, " ")
}

type MigrationLogLogLevel struct {
	value string
}

type MigrationLogLogLevelEnum struct {
	INFO  MigrationLogLogLevel
	ERROR MigrationLogLogLevel
}

func GetMigrationLogLogLevelEnum() MigrationLogLogLevelEnum {
	return MigrationLogLogLevelEnum{
		INFO: MigrationLogLogLevel{
			value: "INFO",
		},
		ERROR: MigrationLogLogLevel{
			value: "ERROR",
		},
	}
}

func (c MigrationLogLogLevel) Value() string {
	return c.value
}

func (c MigrationLogLogLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationLogLogLevel) UnmarshalJSON(b []byte) error {
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
