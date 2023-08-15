package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResponseGroup struct {

	// 规则类别，内置规则(BUILT_IN)或自建规则(BUILT_SELF)
	Category *ResponseGroupCategory `json:"category,omitempty"`

	// 是否允许删除
	DeleteAllowed *bool `json:"delete_allowed,omitempty"`

	// 规则组描述
	GroupDesc *string `json:"group_desc,omitempty"`

	// 规则组名称
	GroupName *string `json:"group_name,omitempty"`

	// 规则组ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	RuleNames *string `json:"rule_names,omitempty"`

	// 扫描任务名称
	TaskNames *string `json:"task_names,omitempty"`

	// 是否为默认规则组
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o ResponseGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseGroup struct{}"
	}

	return strings.Join([]string{"ResponseGroup", string(data)}, " ")
}

type ResponseGroupCategory struct {
	value string
}

type ResponseGroupCategoryEnum struct {
	BUILT_IN   ResponseGroupCategory
	BUILT_SELF ResponseGroupCategory
}

func GetResponseGroupCategoryEnum() ResponseGroupCategoryEnum {
	return ResponseGroupCategoryEnum{
		BUILT_IN: ResponseGroupCategory{
			value: "BUILT_IN",
		},
		BUILT_SELF: ResponseGroupCategory{
			value: "BUILT_SELF",
		},
	}
}

func (c ResponseGroupCategory) Value() string {
	return c.value
}

func (c ResponseGroupCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseGroupCategory) UnmarshalJSON(b []byte) error {
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
