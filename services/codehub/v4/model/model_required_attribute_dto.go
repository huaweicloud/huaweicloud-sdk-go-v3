package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RequiredAttributeDto struct {

	// **参数解释：** 必填字段名称。 描述：Body 严重程度：Severity 指派给：AssigneeId 意见分类：ReviewCategories 意见模块：ReviewModules
	Name *RequiredAttributeDtoName `json:"name,omitempty"`

	// **参数解释：** 是否必填。
	IsRequired *bool `json:"is_required,omitempty"`
}

func (o RequiredAttributeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequiredAttributeDto struct{}"
	}

	return strings.Join([]string{"RequiredAttributeDto", string(data)}, " ")
}

type RequiredAttributeDtoName struct {
	value string
}

type RequiredAttributeDtoNameEnum struct {
	BODY              RequiredAttributeDtoName
	SEVERITY          RequiredAttributeDtoName
	ASSIGNEE_ID       RequiredAttributeDtoName
	REVIEW_CATEGORIES RequiredAttributeDtoName
	REVIEW_MODULES    RequiredAttributeDtoName
}

func GetRequiredAttributeDtoNameEnum() RequiredAttributeDtoNameEnum {
	return RequiredAttributeDtoNameEnum{
		BODY: RequiredAttributeDtoName{
			value: "Body",
		},
		SEVERITY: RequiredAttributeDtoName{
			value: "Severity",
		},
		ASSIGNEE_ID: RequiredAttributeDtoName{
			value: "AssigneeId",
		},
		REVIEW_CATEGORIES: RequiredAttributeDtoName{
			value: "ReviewCategories",
		},
		REVIEW_MODULES: RequiredAttributeDtoName{
			value: "ReviewModules",
		},
	}
}

func (c RequiredAttributeDtoName) Value() string {
	return c.value
}

func (c RequiredAttributeDtoName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RequiredAttributeDtoName) UnmarshalJSON(b []byte) error {
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
