package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchShowParamsRequest struct {
	// 请求语言类型

	XLanguage BatchShowParamsRequestXLanguage `json:"X-Language"`

	Body *BatchQueryParamReq `json:"body,omitempty"`
}

func (o BatchShowParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowParamsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowParamsRequest", string(data)}, " ")
}

type BatchShowParamsRequestXLanguage struct {
	value string
}

type BatchShowParamsRequestXLanguageEnum struct {
	EN_US BatchShowParamsRequestXLanguage
	ZH_CN BatchShowParamsRequestXLanguage
}

func GetBatchShowParamsRequestXLanguageEnum() BatchShowParamsRequestXLanguageEnum {
	return BatchShowParamsRequestXLanguageEnum{
		EN_US: BatchShowParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchShowParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchShowParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
