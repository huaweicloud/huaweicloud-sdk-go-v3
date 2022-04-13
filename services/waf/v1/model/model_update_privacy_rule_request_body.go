package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type UpdatePrivacyRuleRequestBody struct {
	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀

	Url string `json:"url"`
	// 屏蔽字段

	Category UpdatePrivacyRuleRequestBodyCategory `json:"category"`
	// 屏蔽字段名

	Index string `json:"index"`
	// 规则描述

	Description *string `json:"description,omitempty"`
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
