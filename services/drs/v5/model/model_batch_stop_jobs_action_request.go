package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchStopJobsActionRequest Request Object
type BatchStopJobsActionRequest struct {

	// 请求语言类型。
	XLanguage *BatchStopJobsActionRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchStopJobActionReq `json:"body,omitempty"`
}

func (o BatchStopJobsActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobsActionRequest struct{}"
	}

	return strings.Join([]string{"BatchStopJobsActionRequest", string(data)}, " ")
}

type BatchStopJobsActionRequestXLanguage struct {
	value string
}

type BatchStopJobsActionRequestXLanguageEnum struct {
	EN_US BatchStopJobsActionRequestXLanguage
	ZH_CN BatchStopJobsActionRequestXLanguage
}

func GetBatchStopJobsActionRequestXLanguageEnum() BatchStopJobsActionRequestXLanguageEnum {
	return BatchStopJobsActionRequestXLanguageEnum{
		EN_US: BatchStopJobsActionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchStopJobsActionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchStopJobsActionRequestXLanguage) Value() string {
	return c.value
}

func (c BatchStopJobsActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStopJobsActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
