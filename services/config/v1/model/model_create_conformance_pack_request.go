package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConformancePackRequest Request Object
type CreateConformancePackRequest struct {

	// 合规包信息语言，默认为\"en-us\"英文
	XLanguage *CreateConformancePackRequestXLanguage `json:"X-Language,omitempty"`

	Body *ConformancePackRequestBody `json:"body,omitempty"`
}

func (o CreateConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConformancePackRequest struct{}"
	}

	return strings.Join([]string{"CreateConformancePackRequest", string(data)}, " ")
}

type CreateConformancePackRequestXLanguage struct {
	value string
}

type CreateConformancePackRequestXLanguageEnum struct {
	ZH_CN CreateConformancePackRequestXLanguage
	EN_US CreateConformancePackRequestXLanguage
}

func GetCreateConformancePackRequestXLanguageEnum() CreateConformancePackRequestXLanguageEnum {
	return CreateConformancePackRequestXLanguageEnum{
		ZH_CN: CreateConformancePackRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateConformancePackRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateConformancePackRequestXLanguage) Value() string {
	return c.value
}

func (c CreateConformancePackRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConformancePackRequestXLanguage) UnmarshalJSON(b []byte) error {
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
