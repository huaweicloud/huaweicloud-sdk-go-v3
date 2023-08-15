package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteJobsRequest Request Object
type BatchDeleteJobsRequest struct {

	// 请求语言类型
	XLanguage *BatchDeleteJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDeleteJobReq `json:"body,omitempty"`
}

func (o BatchDeleteJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsRequest", string(data)}, " ")
}

type BatchDeleteJobsRequestXLanguage struct {
	value string
}

type BatchDeleteJobsRequestXLanguageEnum struct {
	EN_US BatchDeleteJobsRequestXLanguage
	ZH_CN BatchDeleteJobsRequestXLanguage
}

func GetBatchDeleteJobsRequestXLanguageEnum() BatchDeleteJobsRequestXLanguageEnum {
	return BatchDeleteJobsRequestXLanguageEnum{
		EN_US: BatchDeleteJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchDeleteJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchDeleteJobsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchDeleteJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
