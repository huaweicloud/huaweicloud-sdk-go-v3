package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListRposAndRtosRequest struct {
	// 请求语言类型

	XLanguage BatchListRposAndRtosRequestXLanguage `json:"X-Language"`

	Body *BatchQueryRpoAndRtoReq `json:"body,omitempty"`
}

func (o BatchListRposAndRtosRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosRequest struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosRequest", string(data)}, " ")
}

type BatchListRposAndRtosRequestXLanguage struct {
	value string
}

type BatchListRposAndRtosRequestXLanguageEnum struct {
	EN_US BatchListRposAndRtosRequestXLanguage
	ZH_CN BatchListRposAndRtosRequestXLanguage
}

func GetBatchListRposAndRtosRequestXLanguageEnum() BatchListRposAndRtosRequestXLanguageEnum {
	return BatchListRposAndRtosRequestXLanguageEnum{
		EN_US: BatchListRposAndRtosRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListRposAndRtosRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListRposAndRtosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchListRposAndRtosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
