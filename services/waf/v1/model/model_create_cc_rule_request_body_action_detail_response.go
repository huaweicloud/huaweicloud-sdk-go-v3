package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCcRuleRequestBodyActionDetailResponse 返回页面
type CreateCcRuleRequestBodyActionDetailResponse struct {

	// 内容类型，值可为“application/json”、“text/html”、“text/xml”。
	ContentType *CreateCcRuleRequestBodyActionDetailResponseContentType `json:"content_type,omitempty"`

	// 防护页面内容
	Content *string `json:"content,omitempty"`
}

func (o CreateCcRuleRequestBodyActionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCcRuleRequestBodyActionDetailResponse struct{}"
	}

	return strings.Join([]string{"CreateCcRuleRequestBodyActionDetailResponse", string(data)}, " ")
}

type CreateCcRuleRequestBodyActionDetailResponseContentType struct {
	value string
}

type CreateCcRuleRequestBodyActionDetailResponseContentTypeEnum struct {
	APPLICATION_JSON CreateCcRuleRequestBodyActionDetailResponseContentType
	TEXT_HTML        CreateCcRuleRequestBodyActionDetailResponseContentType
	TEXT_XML         CreateCcRuleRequestBodyActionDetailResponseContentType
}

func GetCreateCcRuleRequestBodyActionDetailResponseContentTypeEnum() CreateCcRuleRequestBodyActionDetailResponseContentTypeEnum {
	return CreateCcRuleRequestBodyActionDetailResponseContentTypeEnum{
		APPLICATION_JSON: CreateCcRuleRequestBodyActionDetailResponseContentType{
			value: "application/json",
		},
		TEXT_HTML: CreateCcRuleRequestBodyActionDetailResponseContentType{
			value: "text/html",
		},
		TEXT_XML: CreateCcRuleRequestBodyActionDetailResponseContentType{
			value: "text/xml",
		},
	}
}

func (c CreateCcRuleRequestBodyActionDetailResponseContentType) Value() string {
	return c.value
}

func (c CreateCcRuleRequestBodyActionDetailResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCcRuleRequestBodyActionDetailResponseContentType) UnmarshalJSON(b []byte) error {
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
