package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAsyncJobsRequest Request Object
type ListAsyncJobsRequest struct {

	// 请求语言类型。
	XLanguage *ListAsyncJobsRequestXLanguage `json:"X-Language,omitempty"`

	// 批量异步创建的任务ID。
	AsyncJobId *string `json:"async_job_id,omitempty"`

	// 批量异步创建的任务状态。
	Status *string `json:"status,omitempty"`

	// 批量异步创建的任务的租户名。
	DomainName *string `json:"domain_name,omitempty"`

	// 批量异步创建的任务的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 返回结果按该关键字排序，默认为“create_time”。
	SortKey *string `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAsyncJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncJobsRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncJobsRequest", string(data)}, " ")
}

type ListAsyncJobsRequestXLanguage struct {
	value string
}

type ListAsyncJobsRequestXLanguageEnum struct {
	EN_US ListAsyncJobsRequestXLanguage
	ZH_CN ListAsyncJobsRequestXLanguage
}

func GetListAsyncJobsRequestXLanguageEnum() ListAsyncJobsRequestXLanguageEnum {
	return ListAsyncJobsRequestXLanguageEnum{
		EN_US: ListAsyncJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAsyncJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAsyncJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ListAsyncJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAsyncJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
