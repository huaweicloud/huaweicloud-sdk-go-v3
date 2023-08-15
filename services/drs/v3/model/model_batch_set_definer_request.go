package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchSetDefinerRequest Request Object
type BatchSetDefinerRequest struct {

	// 请求语言类型
	XLanguage *BatchSetDefinerRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchReplaceDefinerReq `json:"body,omitempty"`
}

func (o BatchSetDefinerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetDefinerRequest struct{}"
	}

	return strings.Join([]string{"BatchSetDefinerRequest", string(data)}, " ")
}

type BatchSetDefinerRequestXLanguage struct {
	value string
}

type BatchSetDefinerRequestXLanguageEnum struct {
	EN_US BatchSetDefinerRequestXLanguage
	ZH_CN BatchSetDefinerRequestXLanguage
}

func GetBatchSetDefinerRequestXLanguageEnum() BatchSetDefinerRequestXLanguageEnum {
	return BatchSetDefinerRequestXLanguageEnum{
		EN_US: BatchSetDefinerRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetDefinerRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetDefinerRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetDefinerRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetDefinerRequestXLanguage) UnmarshalJSON(b []byte) error {
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
