package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchValidateConnectionsRequest struct {
	// 请求语言类型

	XLanguage BatchValidateConnectionsRequestXLanguage `json:"X-Language"`

	Body *BatchTestConnectionReq `json:"body,omitempty"`
}

func (o BatchValidateConnectionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchValidateConnectionsRequest struct{}"
	}

	return strings.Join([]string{"BatchValidateConnectionsRequest", string(data)}, " ")
}

type BatchValidateConnectionsRequestXLanguage struct {
	value string
}

type BatchValidateConnectionsRequestXLanguageEnum struct {
	EN_US BatchValidateConnectionsRequestXLanguage
	ZH_CN BatchValidateConnectionsRequestXLanguage
}

func GetBatchValidateConnectionsRequestXLanguageEnum() BatchValidateConnectionsRequestXLanguageEnum {
	return BatchValidateConnectionsRequestXLanguageEnum{
		EN_US: BatchValidateConnectionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchValidateConnectionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchValidateConnectionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchValidateConnectionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
