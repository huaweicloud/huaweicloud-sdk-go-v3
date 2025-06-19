package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlJobSystemDefendRuleResponse Response Object
type ShowSqlJobSystemDefendRuleResponse struct {

	// 规则类型
	RuleId *ShowSqlJobSystemDefendRuleResponseRuleId `json:"rule_id,omitempty"`

	// 规则状态类型
	Category *ShowSqlJobSystemDefendRuleResponseCategory `json:"category,omitempty"`

	// 可执行的动作
	Actions *[]string `json:"actions,omitempty"`

	// 支持的引擎
	Engines *[]string `json:"engines,omitempty"`

	// 规则是否有限制值
	NoLimit *bool `json:"no_limit,omitempty"`

	// 规则描述
	Desc *string `json:"desc,omitempty"`

	Param          *SysRuleParam `json:"param,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSqlJobSystemDefendRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobSystemDefendRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobSystemDefendRuleResponse", string(data)}, " ")
}

type ShowSqlJobSystemDefendRuleResponseRuleId struct {
	value string
}

type ShowSqlJobSystemDefendRuleResponseRuleIdEnum struct {
	STATIC_0001  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0002  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0003  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0004  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0005  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0006  ShowSqlJobSystemDefendRuleResponseRuleId
	STATIC_0007  ShowSqlJobSystemDefendRuleResponseRuleId
	DYNAMIC_0001 ShowSqlJobSystemDefendRuleResponseRuleId
	DYNAMIC_0002 ShowSqlJobSystemDefendRuleResponseRuleId
	RUNNING_0002 ShowSqlJobSystemDefendRuleResponseRuleId
	RUNNING_0003 ShowSqlJobSystemDefendRuleResponseRuleId
	RUNNING_0004 ShowSqlJobSystemDefendRuleResponseRuleId
}

func GetShowSqlJobSystemDefendRuleResponseRuleIdEnum() ShowSqlJobSystemDefendRuleResponseRuleIdEnum {
	return ShowSqlJobSystemDefendRuleResponseRuleIdEnum{
		STATIC_0001: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0001",
		},
		STATIC_0002: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0002",
		},
		STATIC_0003: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0003",
		},
		STATIC_0004: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0004",
		},
		STATIC_0005: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0005",
		},
		STATIC_0006: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0006",
		},
		STATIC_0007: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "static_0007",
		},
		DYNAMIC_0001: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "dynamic_0001",
		},
		DYNAMIC_0002: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "dynamic_0002",
		},
		RUNNING_0002: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "running_0002",
		},
		RUNNING_0003: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "running_0003",
		},
		RUNNING_0004: ShowSqlJobSystemDefendRuleResponseRuleId{
			value: "running_0004",
		},
	}
}

func (c ShowSqlJobSystemDefendRuleResponseRuleId) Value() string {
	return c.value
}

func (c ShowSqlJobSystemDefendRuleResponseRuleId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobSystemDefendRuleResponseRuleId) UnmarshalJSON(b []byte) error {
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

type ShowSqlJobSystemDefendRuleResponseCategory struct {
	value string
}

type ShowSqlJobSystemDefendRuleResponseCategoryEnum struct {
	STATIC  ShowSqlJobSystemDefendRuleResponseCategory
	DYNAMIC ShowSqlJobSystemDefendRuleResponseCategory
	RUNNING ShowSqlJobSystemDefendRuleResponseCategory
}

func GetShowSqlJobSystemDefendRuleResponseCategoryEnum() ShowSqlJobSystemDefendRuleResponseCategoryEnum {
	return ShowSqlJobSystemDefendRuleResponseCategoryEnum{
		STATIC: ShowSqlJobSystemDefendRuleResponseCategory{
			value: "static",
		},
		DYNAMIC: ShowSqlJobSystemDefendRuleResponseCategory{
			value: "dynamic",
		},
		RUNNING: ShowSqlJobSystemDefendRuleResponseCategory{
			value: "running",
		},
	}
}

func (c ShowSqlJobSystemDefendRuleResponseCategory) Value() string {
	return c.value
}

func (c ShowSqlJobSystemDefendRuleResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobSystemDefendRuleResponseCategory) UnmarshalJSON(b []byte) error {
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
