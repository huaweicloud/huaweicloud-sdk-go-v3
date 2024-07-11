package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAntileakageRuleRequestBody struct {

	// 规则应用的url
	Url string `json:"url"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category CreateAntileakageRuleRequestBodyCategory `json:"category"`

	// 规则内容（http状态码：400 、401、402 、 403 、404 、 405 、500 、501 、502 、503、 504 、507；手机：phone、身份证号：id_card、邮箱：email）
	Contents []string `json:"contents"`

	Action *CreateAntileakageRuleRequestBodyAction `json:"action,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o CreateAntileakageRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntileakageRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAntileakageRuleRequestBody", string(data)}, " ")
}

type CreateAntileakageRuleRequestBodyCategory struct {
	value string
}

type CreateAntileakageRuleRequestBodyCategoryEnum struct {
	CODE      CreateAntileakageRuleRequestBodyCategory
	SENSITIVE CreateAntileakageRuleRequestBodyCategory
}

func GetCreateAntileakageRuleRequestBodyCategoryEnum() CreateAntileakageRuleRequestBodyCategoryEnum {
	return CreateAntileakageRuleRequestBodyCategoryEnum{
		CODE: CreateAntileakageRuleRequestBodyCategory{
			value: "code",
		},
		SENSITIVE: CreateAntileakageRuleRequestBodyCategory{
			value: "sensitive",
		},
	}
}

func (c CreateAntileakageRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c CreateAntileakageRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAntileakageRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
