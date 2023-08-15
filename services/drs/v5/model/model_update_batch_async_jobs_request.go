package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateBatchAsyncJobsRequest Request Object
type UpdateBatchAsyncJobsRequest struct {

	// 批量异步创建的任务ID，由创建批量异步任务接口返回。
	AsyncJobId string `json:"async_job_id"`

	// 请求语言类型。
	XLanguage *UpdateBatchAsyncJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchAsyncUpdateJobReq `json:"body,omitempty"`
}

func (o UpdateBatchAsyncJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchAsyncJobsRequest struct{}"
	}

	return strings.Join([]string{"UpdateBatchAsyncJobsRequest", string(data)}, " ")
}

type UpdateBatchAsyncJobsRequestXLanguage struct {
	value string
}

type UpdateBatchAsyncJobsRequestXLanguageEnum struct {
	EN_US UpdateBatchAsyncJobsRequestXLanguage
	ZH_CN UpdateBatchAsyncJobsRequestXLanguage
}

func GetUpdateBatchAsyncJobsRequestXLanguageEnum() UpdateBatchAsyncJobsRequestXLanguageEnum {
	return UpdateBatchAsyncJobsRequestXLanguageEnum{
		EN_US: UpdateBatchAsyncJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateBatchAsyncJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateBatchAsyncJobsRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateBatchAsyncJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBatchAsyncJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
