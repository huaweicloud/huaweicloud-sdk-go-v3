package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePrivacyRuleRequestBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 隐私屏蔽规则应用的url

	Url *string `json:"url,omitempty"`
	// 屏蔽字段

	Category *CreatePrivacyRuleRequestBodyCategory `json:"category,omitempty"`
	// 屏蔽字段名

	Index *string `json:"index,omitempty"`
}

func (o CreatePrivacyRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivacyRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivacyRuleRequestBody", string(data)}, " ")
}

type CreatePrivacyRuleRequestBodyCategory struct {
	value string
}

type CreatePrivacyRuleRequestBodyCategoryEnum struct {
	PARAMS CreatePrivacyRuleRequestBodyCategory
	COOKIE CreatePrivacyRuleRequestBodyCategory
	HEADER CreatePrivacyRuleRequestBodyCategory
	FORM   CreatePrivacyRuleRequestBodyCategory
}

func GetCreatePrivacyRuleRequestBodyCategoryEnum() CreatePrivacyRuleRequestBodyCategoryEnum {
	return CreatePrivacyRuleRequestBodyCategoryEnum{
		PARAMS: CreatePrivacyRuleRequestBodyCategory{
			value: "params",
		},
		COOKIE: CreatePrivacyRuleRequestBodyCategory{
			value: "cookie",
		},
		HEADER: CreatePrivacyRuleRequestBodyCategory{
			value: "header",
		},
		FORM: CreatePrivacyRuleRequestBodyCategory{
			value: "form",
		},
	}
}

func (c CreatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
