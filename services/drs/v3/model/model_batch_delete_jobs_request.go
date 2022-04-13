package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchDeleteJobsRequest struct {
	// 请求语言类型

	XLanguage BatchDeleteJobsRequestXLanguage `json:"X-Language"`

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

func (c BatchDeleteJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
