package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchValidateClustersConnectionsRequest Request Object
type BatchValidateClustersConnectionsRequest struct {

	// 请求语言类型
	XLanguage *BatchValidateClustersConnectionsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchSpecialTestConnectionReq `json:"body,omitempty"`
}

func (o BatchValidateClustersConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateClustersConnectionsRequest struct{}"
	}

	return strings.Join([]string{"BatchValidateClustersConnectionsRequest", string(data)}, " ")
}

type BatchValidateClustersConnectionsRequestXLanguage struct {
	value string
}

type BatchValidateClustersConnectionsRequestXLanguageEnum struct {
	EN_US BatchValidateClustersConnectionsRequestXLanguage
	ZH_CN BatchValidateClustersConnectionsRequestXLanguage
}

func GetBatchValidateClustersConnectionsRequestXLanguageEnum() BatchValidateClustersConnectionsRequestXLanguageEnum {
	return BatchValidateClustersConnectionsRequestXLanguageEnum{
		EN_US: BatchValidateClustersConnectionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchValidateClustersConnectionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchValidateClustersConnectionsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchValidateClustersConnectionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchValidateClustersConnectionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
