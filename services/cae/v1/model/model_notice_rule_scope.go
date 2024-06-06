package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NoticeRuleScope struct {

	// 生效范围的类型。包括environments（对指定环境下所有组件生效），applications（对指定应用下所有组件生效），components（对指定的组件生效）。
	Type NoticeRuleScopeType `json:"type"`

	// 生效的环境id列表。
	Environments *[]string `json:"environments,omitempty"`

	// 生效的应用id列表。
	Applications *[]string `json:"applications,omitempty"`

	// 生效的组件id列表。
	Components *[]string `json:"components,omitempty"`
}

func (o NoticeRuleScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoticeRuleScope struct{}"
	}

	return strings.Join([]string{"NoticeRuleScope", string(data)}, " ")
}

type NoticeRuleScopeType struct {
	value string
}

type NoticeRuleScopeTypeEnum struct {
	ENVIRONMENTS NoticeRuleScopeType
	APPLICATIONS NoticeRuleScopeType
	COMPONENTS   NoticeRuleScopeType
}

func GetNoticeRuleScopeTypeEnum() NoticeRuleScopeTypeEnum {
	return NoticeRuleScopeTypeEnum{
		ENVIRONMENTS: NoticeRuleScopeType{
			value: "environments",
		},
		APPLICATIONS: NoticeRuleScopeType{
			value: "applications",
		},
		COMPONENTS: NoticeRuleScopeType{
			value: "components",
		},
	}
}

func (c NoticeRuleScopeType) Value() string {
	return c.value
}

func (c NoticeRuleScopeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoticeRuleScopeType) UnmarshalJSON(b []byte) error {
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
