package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSqlLimitRuleOption struct {

	// SQL类型
	SqlType CreateSqlLimitRuleOptionSqlType `json:"sql_type"`

	// 最大并发数
	MaxConcurrency int32 `json:"max_concurrency"`

	// SQL限流规则。限流规则以~分隔关键字，例如select~a。规则举例详细说明：例如关键字是\"select~a\", 含义为：select以及a为该并发控制所包含的两个关键字，~为关键字间隔符，即若执行SQL命令包含select与a两个关键字视为命中此条并发控制规则。
	Pattern string `json:"pattern"`
}

func (o CreateSqlLimitRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitRuleOption struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitRuleOption", string(data)}, " ")
}

type CreateSqlLimitRuleOptionSqlType struct {
	value string
}

type CreateSqlLimitRuleOptionSqlTypeEnum struct {
	SELECT CreateSqlLimitRuleOptionSqlType
	UPDATE CreateSqlLimitRuleOptionSqlType
	DELETE CreateSqlLimitRuleOptionSqlType
}

func GetCreateSqlLimitRuleOptionSqlTypeEnum() CreateSqlLimitRuleOptionSqlTypeEnum {
	return CreateSqlLimitRuleOptionSqlTypeEnum{
		SELECT: CreateSqlLimitRuleOptionSqlType{
			value: "SELECT",
		},
		UPDATE: CreateSqlLimitRuleOptionSqlType{
			value: "UPDATE",
		},
		DELETE: CreateSqlLimitRuleOptionSqlType{
			value: "DELETE",
		},
	}
}

func (c CreateSqlLimitRuleOptionSqlType) Value() string {
	return c.value
}

func (c CreateSqlLimitRuleOptionSqlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlLimitRuleOptionSqlType) UnmarshalJSON(b []byte) error {
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
