package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdatePrivacyRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 隐私屏蔽规则应用的url

	Url *string `json:"url,omitempty"`
	// 屏蔽字段

	Category *UpdatePrivacyRuleResponseCategory `json:"category,omitempty"`
	// 屏蔽字段名

	Index          *string `json:"index,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrivacyRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivacyRuleResponse", string(data)}, " ")
}

type UpdatePrivacyRuleResponseCategory struct {
	value string
}

type UpdatePrivacyRuleResponseCategoryEnum struct {
	PARAMS UpdatePrivacyRuleResponseCategory
	COOKIE UpdatePrivacyRuleResponseCategory
	HEADER UpdatePrivacyRuleResponseCategory
	FORM   UpdatePrivacyRuleResponseCategory
}

func GetUpdatePrivacyRuleResponseCategoryEnum() UpdatePrivacyRuleResponseCategoryEnum {
	return UpdatePrivacyRuleResponseCategoryEnum{
		PARAMS: UpdatePrivacyRuleResponseCategory{
			value: "params",
		},
		COOKIE: UpdatePrivacyRuleResponseCategory{
			value: "cookie",
		},
		HEADER: UpdatePrivacyRuleResponseCategory{
			value: "header",
		},
		FORM: UpdatePrivacyRuleResponseCategory{
			value: "form",
		},
	}
}

func (c UpdatePrivacyRuleResponseCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdatePrivacyRuleResponseCategory) UnmarshalJSON(b []byte) error {
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
