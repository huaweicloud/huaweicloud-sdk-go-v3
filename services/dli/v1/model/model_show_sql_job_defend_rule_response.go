package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlJobDefendRuleResponse Response Object
type ShowSqlJobDefendRuleResponse struct {

	// 规则唯一标识
	RuleUuid *string `json:"rule_uuid,omitempty"`

	// 项目编号
	ProjectId *string `json:"project_id,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则类型
	RuleId *ShowSqlJobDefendRuleResponseRuleId `json:"rule_id,omitempty"`

	// 规则状态类型
	Category *ShowSqlJobDefendRuleResponseCategory `json:"category,omitempty"`

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
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlJobDefendRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDefendRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDefendRuleResponse", string(data)}, " ")
}

type ShowSqlJobDefendRuleResponseRuleId struct {
	value string
}

type ShowSqlJobDefendRuleResponseRuleIdEnum struct {
	STATIC_0001  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0002  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0003  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0004  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0005  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0006  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0007  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0008  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0009  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0010  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0011  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0012  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0013  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0014  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0015  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0016  ShowSqlJobDefendRuleResponseRuleId
	STATIC_0017  ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0001 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0002 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0003 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0004 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0005 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0006 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0007 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0008 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0009 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0010 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0011 ShowSqlJobDefendRuleResponseRuleId
	DYNAMIC_0012 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0001 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0002 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0003 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0004 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0005 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0006 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0007 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0008 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0009 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0010 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0011 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0012 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0013 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0014 ShowSqlJobDefendRuleResponseRuleId
	RUNNING_0015 ShowSqlJobDefendRuleResponseRuleId
}

func GetShowSqlJobDefendRuleResponseRuleIdEnum() ShowSqlJobDefendRuleResponseRuleIdEnum {
	return ShowSqlJobDefendRuleResponseRuleIdEnum{
		STATIC_0001: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0001",
		},
		STATIC_0002: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0002",
		},
		STATIC_0003: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0003",
		},
		STATIC_0004: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0004",
		},
		STATIC_0005: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0005",
		},
		STATIC_0006: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0006",
		},
		STATIC_0007: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0007",
		},
		STATIC_0008: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0008",
		},
		STATIC_0009: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0009",
		},
		STATIC_0010: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0010",
		},
		STATIC_0011: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0011",
		},
		STATIC_0012: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0012",
		},
		STATIC_0013: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0013",
		},
		STATIC_0014: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0014",
		},
		STATIC_0015: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0015",
		},
		STATIC_0016: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0016",
		},
		STATIC_0017: ShowSqlJobDefendRuleResponseRuleId{
			value: "static_0017",
		},
		DYNAMIC_0001: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0001",
		},
		DYNAMIC_0002: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0002",
		},
		DYNAMIC_0003: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0003",
		},
		DYNAMIC_0004: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0004",
		},
		DYNAMIC_0005: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0005",
		},
		DYNAMIC_0006: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0006",
		},
		DYNAMIC_0007: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0007",
		},
		DYNAMIC_0008: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0008",
		},
		DYNAMIC_0009: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0009",
		},
		DYNAMIC_0010: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0010",
		},
		DYNAMIC_0011: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0011",
		},
		DYNAMIC_0012: ShowSqlJobDefendRuleResponseRuleId{
			value: "dynamic_0012",
		},
		RUNNING_0001: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0001",
		},
		RUNNING_0002: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0002",
		},
		RUNNING_0003: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0003",
		},
		RUNNING_0004: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0004",
		},
		RUNNING_0005: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0005",
		},
		RUNNING_0006: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0006",
		},
		RUNNING_0007: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0007",
		},
		RUNNING_0008: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0008",
		},
		RUNNING_0009: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0009",
		},
		RUNNING_0010: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0010",
		},
		RUNNING_0011: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0011",
		},
		RUNNING_0012: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0012",
		},
		RUNNING_0013: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0013",
		},
		RUNNING_0014: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0014",
		},
		RUNNING_0015: ShowSqlJobDefendRuleResponseRuleId{
			value: "running_0015",
		},
	}
}

func (c ShowSqlJobDefendRuleResponseRuleId) Value() string {
	return c.value
}

func (c ShowSqlJobDefendRuleResponseRuleId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobDefendRuleResponseRuleId) UnmarshalJSON(b []byte) error {
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

type ShowSqlJobDefendRuleResponseCategory struct {
	value string
}

type ShowSqlJobDefendRuleResponseCategoryEnum struct {
	STATIC  ShowSqlJobDefendRuleResponseCategory
	DYNAMIC ShowSqlJobDefendRuleResponseCategory
	RUNNING ShowSqlJobDefendRuleResponseCategory
}

func GetShowSqlJobDefendRuleResponseCategoryEnum() ShowSqlJobDefendRuleResponseCategoryEnum {
	return ShowSqlJobDefendRuleResponseCategoryEnum{
		STATIC: ShowSqlJobDefendRuleResponseCategory{
			value: "static",
		},
		DYNAMIC: ShowSqlJobDefendRuleResponseCategory{
			value: "dynamic",
		},
		RUNNING: ShowSqlJobDefendRuleResponseCategory{
			value: "running",
		},
	}
}

func (c ShowSqlJobDefendRuleResponseCategory) Value() string {
	return c.value
}

func (c ShowSqlJobDefendRuleResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobDefendRuleResponseCategory) UnmarshalJSON(b []byte) error {
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
