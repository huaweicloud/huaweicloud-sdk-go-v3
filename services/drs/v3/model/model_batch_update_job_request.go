package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateJobRequest Request Object
type BatchUpdateJobRequest struct {

	// 请求语言类型
	XLanguage *BatchUpdateJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchModifyJobReq `json:"body,omitempty"`
}

func (o BatchUpdateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateJobRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobRequest", string(data)}, " ")
}

type BatchUpdateJobRequestXLanguage struct {
	value string
}

type BatchUpdateJobRequestXLanguageEnum struct {
	EN_US BatchUpdateJobRequestXLanguage
	ZH_CN BatchUpdateJobRequestXLanguage
}

func GetBatchUpdateJobRequestXLanguageEnum() BatchUpdateJobRequestXLanguageEnum {
	return BatchUpdateJobRequestXLanguageEnum{
		EN_US: BatchUpdateJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchUpdateJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchUpdateJobRequestXLanguage) Value() string {
	return c.value
}

func (c BatchUpdateJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
