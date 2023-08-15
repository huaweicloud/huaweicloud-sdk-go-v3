package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchValidateConnectionsRequest Request Object
type BatchValidateConnectionsRequest struct {

	// 请求语言类型
	XLanguage *BatchValidateConnectionsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchTestConnectionReq `json:"body,omitempty"`
}

func (o BatchValidateConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
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

func (c BatchValidateConnectionsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchValidateConnectionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchValidateConnectionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
