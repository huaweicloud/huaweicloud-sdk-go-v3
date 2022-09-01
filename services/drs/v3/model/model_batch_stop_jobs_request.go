package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchStopJobsRequest struct {

	// 请求语言类型
	XLanguage *BatchStopJobsRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	Body *BatchPauseJobReq `json:"body,omitempty" xml:"body"`
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
