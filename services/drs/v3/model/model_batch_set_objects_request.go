package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchSetObjectsRequest Request Object
type BatchSetObjectsRequest struct {

	// 请求语言类型
	XLanguage *BatchSetObjectsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchUpdateDatabaseObjectReq `json:"body,omitempty"`
}

func (o BatchSetObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetObjectsRequest struct{}"
	}

	return strings.Join([]string{"BatchSetObjectsRequest", string(data)}, " ")
}

type BatchSetObjectsRequestXLanguage struct {
	value string
}

type BatchSetObjectsRequestXLanguageEnum struct {
	EN_US BatchSetObjectsRequestXLanguage
	ZH_CN BatchSetObjectsRequestXLanguage
}

func GetBatchSetObjectsRequestXLanguageEnum() BatchSetObjectsRequestXLanguageEnum {
	return BatchSetObjectsRequestXLanguageEnum{
		EN_US: BatchSetObjectsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetObjectsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetObjectsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetObjectsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetObjectsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
