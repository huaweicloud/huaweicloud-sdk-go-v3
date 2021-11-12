package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdatePrivacyRuleRequestBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 隐私屏蔽规则应用的url

	Url *string `json:"url,omitempty"`
	// 屏蔽字段

	Category *UpdatePrivacyRuleRequestBodyCategory `json:"category,omitempty"`
	// 屏蔽字段名

	Index *string `json:"index,omitempty"`
}

func (o UpdatePrivacyRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivacyRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivacyRuleRequestBody", string(data)}, " ")
}

type UpdatePrivacyRuleRequestBodyCategory struct {
	value string
}

type UpdatePrivacyRuleRequestBodyCategoryEnum struct {
	PARAMS UpdatePrivacyRuleRequestBodyCategory
	COOKIE UpdatePrivacyRuleRequestBodyCategory
	HEADER UpdatePrivacyRuleRequestBodyCategory
	FORM   UpdatePrivacyRuleRequestBodyCategory
}

func GetUpdatePrivacyRuleRequestBodyCategoryEnum() UpdatePrivacyRuleRequestBodyCategoryEnum {
	return UpdatePrivacyRuleRequestBodyCategoryEnum{
		PARAMS: UpdatePrivacyRuleRequestBodyCategory{
			value: "params",
		},
		COOKIE: UpdatePrivacyRuleRequestBodyCategory{
			value: "cookie",
		},
		HEADER: UpdatePrivacyRuleRequestBodyCategory{
			value: "header",
		},
		FORM: UpdatePrivacyRuleRequestBodyCategory{
			value: "form",
		},
	}
}

func (c UpdatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
