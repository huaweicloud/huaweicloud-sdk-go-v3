package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListProgressesRequest struct {

	// 请求语言类型
	XLanguage *BatchListProgressesRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryProgressReq `json:"body,omitempty"`
}

func (o BatchListProgressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListProgressesRequest struct{}"
	}

	return strings.Join([]string{"BatchListProgressesRequest", string(data)}, " ")
}

type BatchListProgressesRequestXLanguage struct {
	value string
}

type BatchListProgressesRequestXLanguageEnum struct {
	EN_US BatchListProgressesRequestXLanguage
	ZH_CN BatchListProgressesRequestXLanguage
}

func GetBatchListProgressesRequestXLanguageEnum() BatchListProgressesRequestXLanguageEnum {
	return BatchListProgressesRequestXLanguageEnum{
		EN_US: BatchListProgressesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListProgressesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListProgressesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListProgressesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
