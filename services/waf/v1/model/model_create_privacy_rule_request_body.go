package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePrivacyRuleRequestBody struct {

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"（星号）结尾代表路径前缀
	Url string `json:"url"`

	// 屏蔽字段   - Params：请求参数   - Cookie：根据Cookie区分的Web访问者   - Header：自定义HTTP首部   - Form：表单参数
	Category CreatePrivacyRuleRequestBodyCategory `json:"category"`

	// 屏蔽字段名，根据“屏蔽字段”设置字段名，被屏蔽的字段将不会出现在日志中。屏蔽字段名的长度不能超过2048字节，且只能由数字、字母、下划线和中划线组成。
	Index string `json:"index"`

	// 规则描述，可选参数，设置该规则的备注信息。
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

func (c CreatePrivacyRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c CreatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
