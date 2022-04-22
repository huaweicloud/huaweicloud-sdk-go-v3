package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePrivacyRuleRequestBody struct {

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url string `json:"url"`

	// 屏蔽字段
	Category CreatePrivacyRuleRequestBodyCategory `json:"category"`

	// 屏蔽字段名
	Index string `json:"index"`

	// 规则描述
	Description *string `json:"description,omitempty"`
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
