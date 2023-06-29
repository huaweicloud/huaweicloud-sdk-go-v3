package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CcrulesListInfoActionDetailResponse 阻断页面
type CcrulesListInfoActionDetailResponse struct {

	// 内容类型，值可为“application/json”、“text/html”、“text/xml”。
	ContentType *CcrulesListInfoActionDetailResponseContentType `json:"content_type,omitempty"`

	// 阻断页面内容
	Content *string `json:"content,omitempty"`
}

func (o CcrulesListInfoActionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcrulesListInfoActionDetailResponse struct{}"
	}

	return strings.Join([]string{"CcrulesListInfoActionDetailResponse", string(data)}, " ")
}

type CcrulesListInfoActionDetailResponseContentType struct {
	value string
}

type CcrulesListInfoActionDetailResponseContentTypeEnum struct {
	APPLICATION_JSON CcrulesListInfoActionDetailResponseContentType
	TEXT_HTML        CcrulesListInfoActionDetailResponseContentType
	TEXT_XML         CcrulesListInfoActionDetailResponseContentType
}

func GetCcrulesListInfoActionDetailResponseContentTypeEnum() CcrulesListInfoActionDetailResponseContentTypeEnum {
	return CcrulesListInfoActionDetailResponseContentTypeEnum{
		APPLICATION_JSON: CcrulesListInfoActionDetailResponseContentType{
			value: "application/json",
		},
		TEXT_HTML: CcrulesListInfoActionDetailResponseContentType{
			value: "text/html",
		},
		TEXT_XML: CcrulesListInfoActionDetailResponseContentType{
			value: "text/xml",
		},
	}
}

func (c CcrulesListInfoActionDetailResponseContentType) Value() string {
	return c.value
}

func (c CcrulesListInfoActionDetailResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CcrulesListInfoActionDetailResponseContentType) UnmarshalJSON(b []byte) error {
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
