package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlJobSystemDefendRuleResponseBody 系统预制可创建的sql拦截规则
type ShowSqlJobSystemDefendRuleResponseBody struct {

	// 规则类型
	RuleId *ShowSqlJobSystemDefendRuleResponseBodyRuleId `json:"rule_id,omitempty"`

	// 规则状态类型
	Category *ShowSqlJobSystemDefendRuleResponseBodyCategory `json:"category,omitempty"`

	// 可执行的动作
	Actions *[]string `json:"actions,omitempty"`

	// 支持的引擎
	Engines *[]string `json:"engines,omitempty"`

	// 规则是否有限制值
	NoLimit *bool `json:"no_limit,omitempty"`

	// 规则描述
	Desc *string `json:"desc,omitempty"`

	Param *SysRuleParam `json:"param,omitempty"`
}

func (o ShowSqlJobSystemDefendRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobSystemDefendRuleResponseBody struct{}"
	}

	return strings.Join([]string{"ShowSqlJobSystemDefendRuleResponseBody", string(data)}, " ")
}

type ShowSqlJobSystemDefendRuleResponseBodyRuleId struct {
	value string
}

type ShowSqlJobSystemDefendRuleResponseBodyRuleIdEnum struct {
	STATIC_0001  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0002  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0003  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0004  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0005  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0006  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	STATIC_0007  ShowSqlJobSystemDefendRuleResponseBodyRuleId
	DYNAMIC_0001 ShowSqlJobSystemDefendRuleResponseBodyRuleId
	DYNAMIC_0002 ShowSqlJobSystemDefendRuleResponseBodyRuleId
	RUNNING_0002 ShowSqlJobSystemDefendRuleResponseBodyRuleId
	RUNNING_0003 ShowSqlJobSystemDefendRuleResponseBodyRuleId
	RUNNING_0004 ShowSqlJobSystemDefendRuleResponseBodyRuleId
}

func GetShowSqlJobSystemDefendRuleResponseBodyRuleIdEnum() ShowSqlJobSystemDefendRuleResponseBodyRuleIdEnum {
	return ShowSqlJobSystemDefendRuleResponseBodyRuleIdEnum{
		STATIC_0001: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0001",
		},
		STATIC_0002: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0002",
		},
		STATIC_0003: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0003",
		},
		STATIC_0004: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0004",
		},
		STATIC_0005: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0005",
		},
		STATIC_0006: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0006",
		},
		STATIC_0007: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "static_0007",
		},
		DYNAMIC_0001: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "dynamic_0001",
		},
		DYNAMIC_0002: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "dynamic_0002",
		},
		RUNNING_0002: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "running_0002",
		},
		RUNNING_0003: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "running_0003",
		},
		RUNNING_0004: ShowSqlJobSystemDefendRuleResponseBodyRuleId{
			value: "running_0004",
		},
	}
}

func (c ShowSqlJobSystemDefendRuleResponseBodyRuleId) Value() string {
	return c.value
}

func (c ShowSqlJobSystemDefendRuleResponseBodyRuleId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobSystemDefendRuleResponseBodyRuleId) UnmarshalJSON(b []byte) error {
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

type ShowSqlJobSystemDefendRuleResponseBodyCategory struct {
	value string
}

type ShowSqlJobSystemDefendRuleResponseBodyCategoryEnum struct {
	STATIC  ShowSqlJobSystemDefendRuleResponseBodyCategory
	DYNAMIC ShowSqlJobSystemDefendRuleResponseBodyCategory
	RUNNING ShowSqlJobSystemDefendRuleResponseBodyCategory
}

func GetShowSqlJobSystemDefendRuleResponseBodyCategoryEnum() ShowSqlJobSystemDefendRuleResponseBodyCategoryEnum {
	return ShowSqlJobSystemDefendRuleResponseBodyCategoryEnum{
		STATIC: ShowSqlJobSystemDefendRuleResponseBodyCategory{
			value: "static",
		},
		DYNAMIC: ShowSqlJobSystemDefendRuleResponseBodyCategory{
			value: "dynamic",
		},
		RUNNING: ShowSqlJobSystemDefendRuleResponseBodyCategory{
			value: "running",
		},
	}
}

func (c ShowSqlJobSystemDefendRuleResponseBodyCategory) Value() string {
	return c.value
}

func (c ShowSqlJobSystemDefendRuleResponseBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobSystemDefendRuleResponseBodyCategory) UnmarshalJSON(b []byte) error {
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
