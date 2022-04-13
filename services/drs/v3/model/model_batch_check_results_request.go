package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchCheckResultsRequest struct {
	// 请求语言类型

	XLanguage BatchCheckResultsRequestXLanguage `json:"X-Language"`

	Body *BatchQueryPrecheckResultReq `json:"body,omitempty"`
}

func (o BatchCheckResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckResultsRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckResultsRequest", string(data)}, " ")
}

type BatchCheckResultsRequestXLanguage struct {
	value string
}

type BatchCheckResultsRequestXLanguageEnum struct {
	EN_US BatchCheckResultsRequestXLanguage
	ZH_CN BatchCheckResultsRequestXLanguage
}

func GetBatchCheckResultsRequestXLanguageEnum() BatchCheckResultsRequestXLanguageEnum {
	return BatchCheckResultsRequestXLanguageEnum{
		EN_US: BatchCheckResultsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCheckResultsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCheckResultsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCheckResultsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
