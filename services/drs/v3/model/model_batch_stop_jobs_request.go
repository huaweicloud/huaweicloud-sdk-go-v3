package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchStopJobsRequest Request Object
type BatchStopJobsRequest struct {

	// 请求语言类型
	XLanguage *BatchStopJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchPauseJobReq `json:"body,omitempty"`
}

func (o BatchStopJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchStopJobsRequest", string(data)}, " ")
}

type BatchStopJobsRequestXLanguage struct {
	value string
}

type BatchStopJobsRequestXLanguageEnum struct {
	EN_US BatchStopJobsRequestXLanguage
	ZH_CN BatchStopJobsRequestXLanguage
}

func GetBatchStopJobsRequestXLanguageEnum() BatchStopJobsRequestXLanguageEnum {
	return BatchStopJobsRequestXLanguageEnum{
		EN_US: BatchStopJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchStopJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchStopJobsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchStopJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStopJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
