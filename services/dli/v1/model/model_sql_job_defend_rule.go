package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlJobDefendRule 用户创建的sql拦截规则
type SqlJobDefendRule struct {

	// 规则唯一标识
	RuleUuid *string `json:"rule_uuid,omitempty"`

	// 项目编号
	ProjectId *string `json:"project_id,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则类型
	RuleId *SqlJobDefendRuleRuleId `json:"rule_id,omitempty"`

	// 规则状态类型
	Category *SqlJobDefendRuleCategory `json:"category,omitempty"`

	// 规则详情
	EngineRules *interface{} `json:"engine_rules,omitempty"`

	// 队列名称
	QueueNames *[]string `json:"queueNames,omitempty"`

	// 用户规则描述
	Desc *string `json:"desc,omitempty"`

	// 系统规则描述
	SysDesc *string `json:"sys_desc,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o SqlJobDefendRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobDefendRule struct{}"
	}

	return strings.Join([]string{"SqlJobDefendRule", string(data)}, " ")
}

type SqlJobDefendRuleRuleId struct {
	value string
}

type SqlJobDefendRuleRuleIdEnum struct {
	STATIC_0001  SqlJobDefendRuleRuleId
	STATIC_0002  SqlJobDefendRuleRuleId
	STATIC_0003  SqlJobDefendRuleRuleId
	STATIC_0004  SqlJobDefendRuleRuleId
	STATIC_0005  SqlJobDefendRuleRuleId
	STATIC_0006  SqlJobDefendRuleRuleId
	STATIC_0007  SqlJobDefendRuleRuleId
	STATIC_0008  SqlJobDefendRuleRuleId
	STATIC_0009  SqlJobDefendRuleRuleId
	STATIC_0010  SqlJobDefendRuleRuleId
	STATIC_0011  SqlJobDefendRuleRuleId
	STATIC_0012  SqlJobDefendRuleRuleId
	STATIC_0013  SqlJobDefendRuleRuleId
	STATIC_0014  SqlJobDefendRuleRuleId
	STATIC_0015  SqlJobDefendRuleRuleId
	STATIC_0016  SqlJobDefendRuleRuleId
	STATIC_0017  SqlJobDefendRuleRuleId
	DYNAMIC_0001 SqlJobDefendRuleRuleId
	DYNAMIC_0002 SqlJobDefendRuleRuleId
	DYNAMIC_0003 SqlJobDefendRuleRuleId
	DYNAMIC_0004 SqlJobDefendRuleRuleId
	DYNAMIC_0005 SqlJobDefendRuleRuleId
	DYNAMIC_0006 SqlJobDefendRuleRuleId
	DYNAMIC_0007 SqlJobDefendRuleRuleId
	DYNAMIC_0008 SqlJobDefendRuleRuleId
	DYNAMIC_0009 SqlJobDefendRuleRuleId
	DYNAMIC_0010 SqlJobDefendRuleRuleId
	DYNAMIC_0011 SqlJobDefendRuleRuleId
	DYNAMIC_0012 SqlJobDefendRuleRuleId
	RUNNING_0001 SqlJobDefendRuleRuleId
	RUNNING_0002 SqlJobDefendRuleRuleId
	RUNNING_0003 SqlJobDefendRuleRuleId
	RUNNING_0004 SqlJobDefendRuleRuleId
	RUNNING_0005 SqlJobDefendRuleRuleId
	RUNNING_0006 SqlJobDefendRuleRuleId
	RUNNING_0007 SqlJobDefendRuleRuleId
	RUNNING_0008 SqlJobDefendRuleRuleId
	RUNNING_0009 SqlJobDefendRuleRuleId
	RUNNING_0010 SqlJobDefendRuleRuleId
	RUNNING_0011 SqlJobDefendRuleRuleId
	RUNNING_0012 SqlJobDefendRuleRuleId
	RUNNING_0013 SqlJobDefendRuleRuleId
	RUNNING_0014 SqlJobDefendRuleRuleId
	RUNNING_0015 SqlJobDefendRuleRuleId
}

func GetSqlJobDefendRuleRuleIdEnum() SqlJobDefendRuleRuleIdEnum {
	return SqlJobDefendRuleRuleIdEnum{
		STATIC_0001: SqlJobDefendRuleRuleId{
			value: "static_0001",
		},
		STATIC_0002: SqlJobDefendRuleRuleId{
			value: "static_0002",
		},
		STATIC_0003: SqlJobDefendRuleRuleId{
			value: "static_0003",
		},
		STATIC_0004: SqlJobDefendRuleRuleId{
			value: "static_0004",
		},
		STATIC_0005: SqlJobDefendRuleRuleId{
			value: "static_0005",
		},
		STATIC_0006: SqlJobDefendRuleRuleId{
			value: "static_0006",
		},
		STATIC_0007: SqlJobDefendRuleRuleId{
			value: "static_0007",
		},
		STATIC_0008: SqlJobDefendRuleRuleId{
			value: "static_0008",
		},
		STATIC_0009: SqlJobDefendRuleRuleId{
			value: "static_0009",
		},
		STATIC_0010: SqlJobDefendRuleRuleId{
			value: "static_0010",
		},
		STATIC_0011: SqlJobDefendRuleRuleId{
			value: "static_0011",
		},
		STATIC_0012: SqlJobDefendRuleRuleId{
			value: "static_0012",
		},
		STATIC_0013: SqlJobDefendRuleRuleId{
			value: "static_0013",
		},
		STATIC_0014: SqlJobDefendRuleRuleId{
			value: "static_0014",
		},
		STATIC_0015: SqlJobDefendRuleRuleId{
			value: "static_0015",
		},
		STATIC_0016: SqlJobDefendRuleRuleId{
			value: "static_0016",
		},
		STATIC_0017: SqlJobDefendRuleRuleId{
			value: "static_0017",
		},
		DYNAMIC_0001: SqlJobDefendRuleRuleId{
			value: "dynamic_0001",
		},
		DYNAMIC_0002: SqlJobDefendRuleRuleId{
			value: "dynamic_0002",
		},
		DYNAMIC_0003: SqlJobDefendRuleRuleId{
			value: "dynamic_0003",
		},
		DYNAMIC_0004: SqlJobDefendRuleRuleId{
			value: "dynamic_0004",
		},
		DYNAMIC_0005: SqlJobDefendRuleRuleId{
			value: "dynamic_0005",
		},
		DYNAMIC_0006: SqlJobDefendRuleRuleId{
			value: "dynamic_0006",
		},
		DYNAMIC_0007: SqlJobDefendRuleRuleId{
			value: "dynamic_0007",
		},
		DYNAMIC_0008: SqlJobDefendRuleRuleId{
			value: "dynamic_0008",
		},
		DYNAMIC_0009: SqlJobDefendRuleRuleId{
			value: "dynamic_0009",
		},
		DYNAMIC_0010: SqlJobDefendRuleRuleId{
			value: "dynamic_0010",
		},
		DYNAMIC_0011: SqlJobDefendRuleRuleId{
			value: "dynamic_0011",
		},
		DYNAMIC_0012: SqlJobDefendRuleRuleId{
			value: "dynamic_0012",
		},
		RUNNING_0001: SqlJobDefendRuleRuleId{
			value: "running_0001",
		},
		RUNNING_0002: SqlJobDefendRuleRuleId{
			value: "running_0002",
		},
		RUNNING_0003: SqlJobDefendRuleRuleId{
			value: "running_0003",
		},
		RUNNING_0004: SqlJobDefendRuleRuleId{
			value: "running_0004",
		},
		RUNNING_0005: SqlJobDefendRuleRuleId{
			value: "running_0005",
		},
		RUNNING_0006: SqlJobDefendRuleRuleId{
			value: "running_0006",
		},
		RUNNING_0007: SqlJobDefendRuleRuleId{
			value: "running_0007",
		},
		RUNNING_0008: SqlJobDefendRuleRuleId{
			value: "running_0008",
		},
		RUNNING_0009: SqlJobDefendRuleRuleId{
			value: "running_0009",
		},
		RUNNING_0010: SqlJobDefendRuleRuleId{
			value: "running_0010",
		},
		RUNNING_0011: SqlJobDefendRuleRuleId{
			value: "running_0011",
		},
		RUNNING_0012: SqlJobDefendRuleRuleId{
			value: "running_0012",
		},
		RUNNING_0013: SqlJobDefendRuleRuleId{
			value: "running_0013",
		},
		RUNNING_0014: SqlJobDefendRuleRuleId{
			value: "running_0014",
		},
		RUNNING_0015: SqlJobDefendRuleRuleId{
			value: "running_0015",
		},
	}
}

func (c SqlJobDefendRuleRuleId) Value() string {
	return c.value
}

func (c SqlJobDefendRuleRuleId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobDefendRuleRuleId) UnmarshalJSON(b []byte) error {
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

type SqlJobDefendRuleCategory struct {
	value string
}

type SqlJobDefendRuleCategoryEnum struct {
	STATIC  SqlJobDefendRuleCategory
	DYNAMIC SqlJobDefendRuleCategory
	RUNNING SqlJobDefendRuleCategory
}

func GetSqlJobDefendRuleCategoryEnum() SqlJobDefendRuleCategoryEnum {
	return SqlJobDefendRuleCategoryEnum{
		STATIC: SqlJobDefendRuleCategory{
			value: "static",
		},
		DYNAMIC: SqlJobDefendRuleCategory{
			value: "dynamic",
		},
		RUNNING: SqlJobDefendRuleCategory{
			value: "running",
		},
	}
}

func (c SqlJobDefendRuleCategory) Value() string {
	return c.value
}

func (c SqlJobDefendRuleCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobDefendRuleCategory) UnmarshalJSON(b []byte) error {
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
