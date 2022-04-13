package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CreatePrivacyRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀

	Url *string `json:"url,omitempty"`
	// 屏蔽字段

	Category *CreatePrivacyRuleResponseCategory `json:"category,omitempty"`
	// 屏蔽字段名

	Index          *string `json:"index,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivacyRuleResponse", string(data)}, " ")
}

type CreatePrivacyRuleResponseCategory struct {
	value string
}

type CreatePrivacyRuleResponseCategoryEnum struct {
	PARAMS CreatePrivacyRuleResponseCategory
	COOKIE CreatePrivacyRuleResponseCategory
	HEADER CreatePrivacyRuleResponseCategory
	FORM   CreatePrivacyRuleResponseCategory
}

func GetCreatePrivacyRuleResponseCategoryEnum() CreatePrivacyRuleResponseCategoryEnum {
	return CreatePrivacyRuleResponseCategoryEnum{
		PARAMS: CreatePrivacyRuleResponseCategory{
			value: "params",
		},
		COOKIE: CreatePrivacyRuleResponseCategory{
			value: "cookie",
		},
		HEADER: CreatePrivacyRuleResponseCategory{
			value: "header",
		},
		FORM: CreatePrivacyRuleResponseCategory{
			value: "form",
		},
	}
}

func (c CreatePrivacyRuleResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivacyRuleResponseCategory) UnmarshalJSON(b []byte) error {
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
