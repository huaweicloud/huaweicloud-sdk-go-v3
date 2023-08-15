package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateJobsAsyncRequest Request Object
type BatchCreateJobsAsyncRequest struct {

	// 请求语言类型。
	XLanguage *BatchCreateJobsAsyncRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchAsyncCreateJobReq `json:"body,omitempty"`
}

func (o BatchCreateJobsAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsAsyncRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsAsyncRequest", string(data)}, " ")
}

type BatchCreateJobsAsyncRequestXLanguage struct {
	value string
}

type BatchCreateJobsAsyncRequestXLanguageEnum struct {
	EN_US BatchCreateJobsAsyncRequestXLanguage
	ZH_CN BatchCreateJobsAsyncRequestXLanguage
}

func GetBatchCreateJobsAsyncRequestXLanguageEnum() BatchCreateJobsAsyncRequestXLanguageEnum {
	return BatchCreateJobsAsyncRequestXLanguageEnum{
		EN_US: BatchCreateJobsAsyncRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCreateJobsAsyncRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCreateJobsAsyncRequestXLanguage) Value() string {
	return c.value
}

func (c BatchCreateJobsAsyncRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateJobsAsyncRequestXLanguage) UnmarshalJSON(b []byte) error {
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
