package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchUpdateAntileakageRuleRequestBody struct {

	// 规则应用的url
	Url string `json:"url"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category BatchUpdateAntileakageRuleRequestBodyCategory `json:"category"`

	// 内容（http状态码：400 、401、402 、 403 、404 、 405 、500 、501 、502 、503、 504 、507；手机：phone、身份证号：id_card、邮箱：email）
	Contents []string `json:"contents"`

	Action *BatchCreateAntileakageRuleRequestBodyAction `json:"action,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds *[]PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids,omitempty"`
}

func (o BatchUpdateAntileakageRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAntileakageRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateAntileakageRuleRequestBody", string(data)}, " ")
}

type BatchUpdateAntileakageRuleRequestBodyCategory struct {
	value string
}

type BatchUpdateAntileakageRuleRequestBodyCategoryEnum struct {
	CODE      BatchUpdateAntileakageRuleRequestBodyCategory
	SENSITIVE BatchUpdateAntileakageRuleRequestBodyCategory
}

func GetBatchUpdateAntileakageRuleRequestBodyCategoryEnum() BatchUpdateAntileakageRuleRequestBodyCategoryEnum {
	return BatchUpdateAntileakageRuleRequestBodyCategoryEnum{
		CODE: BatchUpdateAntileakageRuleRequestBodyCategory{
			value: "code",
		},
		SENSITIVE: BatchUpdateAntileakageRuleRequestBodyCategory{
			value: "sensitive",
		},
	}
}

func (c BatchUpdateAntileakageRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c BatchUpdateAntileakageRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateAntileakageRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
