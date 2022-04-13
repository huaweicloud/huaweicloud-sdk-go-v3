package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchSetDefinerRequest struct {
	// 请求语言类型

	XLanguage BatchSetDefinerRequestXLanguage `json:"X-Language"`

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

func (c BatchSetDefinerRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetDefinerRequestXLanguage) UnmarshalJSON(b []byte) error {
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
