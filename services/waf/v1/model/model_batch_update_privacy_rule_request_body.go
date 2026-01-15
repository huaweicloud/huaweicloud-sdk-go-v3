package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchUpdatePrivacyRuleRequestBody struct {

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"（星号）结尾代表路径前缀
	Url string `json:"url"`

	// **参数解释：** 屏蔽字段 **约束限制：** 不涉及 **取值范围：**  - params: 请求参数  - cookie: 根据Cookie区分的Web访问者  - header: 自定义HTTP首部  - form: 表单参数  **默认取值：** 不涉及
	Category BatchUpdatePrivacyRuleRequestBodyCategory `json:"category"`

	// 屏蔽字段名，根据“屏蔽字段”设置字段名，被屏蔽的字段将不会出现在日志中。屏蔽字段名长度不能超过2048字节，且只能由数字、字母、下划线和中划线组成
	Index string `json:"index"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds []PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids"`
}

func (o BatchUpdatePrivacyRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePrivacyRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePrivacyRuleRequestBody", string(data)}, " ")
}

type BatchUpdatePrivacyRuleRequestBodyCategory struct {
	value string
}

type BatchUpdatePrivacyRuleRequestBodyCategoryEnum struct {
	PARAMS BatchUpdatePrivacyRuleRequestBodyCategory
	COOKIE BatchUpdatePrivacyRuleRequestBodyCategory
	HEADER BatchUpdatePrivacyRuleRequestBodyCategory
	FORM   BatchUpdatePrivacyRuleRequestBodyCategory
}

func GetBatchUpdatePrivacyRuleRequestBodyCategoryEnum() BatchUpdatePrivacyRuleRequestBodyCategoryEnum {
	return BatchUpdatePrivacyRuleRequestBodyCategoryEnum{
		PARAMS: BatchUpdatePrivacyRuleRequestBodyCategory{
			value: "params",
		},
		COOKIE: BatchUpdatePrivacyRuleRequestBodyCategory{
			value: "cookie",
		},
		HEADER: BatchUpdatePrivacyRuleRequestBodyCategory{
			value: "header",
		},
		FORM: BatchUpdatePrivacyRuleRequestBodyCategory{
			value: "form",
		},
	}
}

func (c BatchUpdatePrivacyRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c BatchUpdatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
