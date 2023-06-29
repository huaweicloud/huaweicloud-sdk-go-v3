package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePrivacyRuleResponse Response Object
type CreatePrivacyRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间，格式为13位毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url *string `json:"url,omitempty"`

	// 屏蔽字段   - Params：请求参数   - Cookie：根据Cookie区分的Web访问者   - Header：自定义HTTP首部   - Form：表单参数
	Category *CreatePrivacyRuleResponseCategory `json:"category,omitempty"`

	// 屏蔽字段名，根据“屏蔽字段”设置字段名，被屏蔽的字段将不会出现在日志中。
	Index *string `json:"index,omitempty"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description    *string `json:"description,omitempty"`
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

func (c CreatePrivacyRuleResponseCategory) Value() string {
	return c.value
}

func (c CreatePrivacyRuleResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivacyRuleResponseCategory) UnmarshalJSON(b []byte) error {
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
