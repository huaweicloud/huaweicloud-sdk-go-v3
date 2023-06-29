package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RuleGroupRequest struct {

	// 规则类别，内置规则(BUILT_IN)或自建规则(BUILT_SELF)
	Category *RuleGroupRequestCategory `json:"category,omitempty"`

	// 是否默认规则组
	DefaultStatus *bool `json:"default_status,omitempty"`

	// 规则组描述
	GroupDesc *string `json:"group_desc,omitempty"`

	// 规则组名称
	GroupName *string `json:"group_name,omitempty"`

	// 规则组ID
	Id *string `json:"id,omitempty"`

	// 包含的规则ID列表
	RuleIds *[]string `json:"rule_ids,omitempty"`
}

func (o RuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleGroupRequest struct{}"
	}

	return strings.Join([]string{"RuleGroupRequest", string(data)}, " ")
}

type RuleGroupRequestCategory struct {
	value string
}

type RuleGroupRequestCategoryEnum struct {
	BUILT_IN   RuleGroupRequestCategory
	BUILT_SELF RuleGroupRequestCategory
}

func GetRuleGroupRequestCategoryEnum() RuleGroupRequestCategoryEnum {
	return RuleGroupRequestCategoryEnum{
		BUILT_IN: RuleGroupRequestCategory{
			value: "BUILT_IN",
		},
		BUILT_SELF: RuleGroupRequestCategory{
			value: "BUILT_SELF",
		},
	}
}

func (c RuleGroupRequestCategory) Value() string {
	return c.value
}

func (c RuleGroupRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleGroupRequestCategory) UnmarshalJSON(b []byte) error {
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
