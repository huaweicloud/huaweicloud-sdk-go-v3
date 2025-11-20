package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateAntileakageRuleRequestBody struct {

	// 规则应用的url
	Url string `json:"url"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category BatchCreateAntileakageRuleRequestBodyCategory `json:"category"`

	// 规则内容（http状态码：400 、401、402 、 403 、404 、 405 、500 、501 、502 、503、 504 、507；手机：phone、身份证号：id_card、邮箱：email）
	Contents []string `json:"contents"`

	Action *LeakageListInfoAction `json:"action,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateAntileakageRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAntileakageRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateAntileakageRuleRequestBody", string(data)}, " ")
}

type BatchCreateAntileakageRuleRequestBodyCategory struct {
	value string
}

type BatchCreateAntileakageRuleRequestBodyCategoryEnum struct {
	CODE      BatchCreateAntileakageRuleRequestBodyCategory
	SENSITIVE BatchCreateAntileakageRuleRequestBodyCategory
}

func GetBatchCreateAntileakageRuleRequestBodyCategoryEnum() BatchCreateAntileakageRuleRequestBodyCategoryEnum {
	return BatchCreateAntileakageRuleRequestBodyCategoryEnum{
		CODE: BatchCreateAntileakageRuleRequestBodyCategory{
			value: "code",
		},
		SENSITIVE: BatchCreateAntileakageRuleRequestBodyCategory{
			value: "sensitive",
		},
	}
}

func (c BatchCreateAntileakageRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c BatchCreateAntileakageRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateAntileakageRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
