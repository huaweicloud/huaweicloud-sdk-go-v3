package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlJobDefendRuleRequestBody 用户创建的sql拦截规则
type SqlJobDefendRuleRequestBody struct {

	// 规则名称
	RuleName string `json:"rule_name"`

	// 规则类型
	RuleId SqlJobDefendRuleRequestBodyRuleId `json:"rule_id"`

	// 规则状态类型
	Category SqlJobDefendRuleRequestBodyCategory `json:"category"`

	// 规则详情
	EngineRules *interface{} `json:"engine_rules"`

	// 队列名称
	QueueNames *[]string `json:"queueNames,omitempty"`
}

func (o SqlJobDefendRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobDefendRuleRequestBody struct{}"
	}

	return strings.Join([]string{"SqlJobDefendRuleRequestBody", string(data)}, " ")
}

type SqlJobDefendRuleRequestBodyRuleId struct {
	value string
}

type SqlJobDefendRuleRequestBodyRuleIdEnum struct {
	STATIC_0001  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0002  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0003  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0004  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0005  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0006  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0007  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0008  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0009  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0010  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0011  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0012  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0013  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0014  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0015  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0016  SqlJobDefendRuleRequestBodyRuleId
	STATIC_0017  SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0001 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0002 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0003 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0004 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0005 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0006 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0007 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0008 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0009 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0010 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0011 SqlJobDefendRuleRequestBodyRuleId
	DYNAMIC_0012 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0001 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0002 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0003 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0004 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0005 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0006 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0007 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0008 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0009 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0010 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0011 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0012 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0013 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0014 SqlJobDefendRuleRequestBodyRuleId
	RUNNING_0015 SqlJobDefendRuleRequestBodyRuleId
}

func GetSqlJobDefendRuleRequestBodyRuleIdEnum() SqlJobDefendRuleRequestBodyRuleIdEnum {
	return SqlJobDefendRuleRequestBodyRuleIdEnum{
		STATIC_0001: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0001",
		},
		STATIC_0002: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0002",
		},
		STATIC_0003: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0003",
		},
		STATIC_0004: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0004",
		},
		STATIC_0005: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0005",
		},
		STATIC_0006: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0006",
		},
		STATIC_0007: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0007",
		},
		STATIC_0008: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0008",
		},
		STATIC_0009: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0009",
		},
		STATIC_0010: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0010",
		},
		STATIC_0011: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0011",
		},
		STATIC_0012: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0012",
		},
		STATIC_0013: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0013",
		},
		STATIC_0014: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0014",
		},
		STATIC_0015: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0015",
		},
		STATIC_0016: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0016",
		},
		STATIC_0017: SqlJobDefendRuleRequestBodyRuleId{
			value: "static_0017",
		},
		DYNAMIC_0001: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0001",
		},
		DYNAMIC_0002: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0002",
		},
		DYNAMIC_0003: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0003",
		},
		DYNAMIC_0004: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0004",
		},
		DYNAMIC_0005: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0005",
		},
		DYNAMIC_0006: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0006",
		},
		DYNAMIC_0007: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0007",
		},
		DYNAMIC_0008: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0008",
		},
		DYNAMIC_0009: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0009",
		},
		DYNAMIC_0010: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0010",
		},
		DYNAMIC_0011: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0011",
		},
		DYNAMIC_0012: SqlJobDefendRuleRequestBodyRuleId{
			value: "dynamic_0012",
		},
		RUNNING_0001: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0001",
		},
		RUNNING_0002: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0002",
		},
		RUNNING_0003: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0003",
		},
		RUNNING_0004: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0004",
		},
		RUNNING_0005: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0005",
		},
		RUNNING_0006: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0006",
		},
		RUNNING_0007: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0007",
		},
		RUNNING_0008: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0008",
		},
		RUNNING_0009: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0009",
		},
		RUNNING_0010: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0010",
		},
		RUNNING_0011: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0011",
		},
		RUNNING_0012: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0012",
		},
		RUNNING_0013: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0013",
		},
		RUNNING_0014: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0014",
		},
		RUNNING_0015: SqlJobDefendRuleRequestBodyRuleId{
			value: "running_0015",
		},
	}
}

func (c SqlJobDefendRuleRequestBodyRuleId) Value() string {
	return c.value
}

func (c SqlJobDefendRuleRequestBodyRuleId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobDefendRuleRequestBodyRuleId) UnmarshalJSON(b []byte) error {
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

type SqlJobDefendRuleRequestBodyCategory struct {
	value string
}

type SqlJobDefendRuleRequestBodyCategoryEnum struct {
	STATIC  SqlJobDefendRuleRequestBodyCategory
	DYNAMIC SqlJobDefendRuleRequestBodyCategory
	RUNNING SqlJobDefendRuleRequestBodyCategory
}

func GetSqlJobDefendRuleRequestBodyCategoryEnum() SqlJobDefendRuleRequestBodyCategoryEnum {
	return SqlJobDefendRuleRequestBodyCategoryEnum{
		STATIC: SqlJobDefendRuleRequestBodyCategory{
			value: "static",
		},
		DYNAMIC: SqlJobDefendRuleRequestBodyCategory{
			value: "dynamic",
		},
		RUNNING: SqlJobDefendRuleRequestBodyCategory{
			value: "running",
		},
	}
}

func (c SqlJobDefendRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c SqlJobDefendRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobDefendRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
