package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAsyncJobDetailRequest struct {

	// 批量异步创建的任务ID，由创建批量异步任务接口返回。
	AsyncJobId string `json:"async_job_id"`

	// 请求语言类型。
	XLanguage *ListAsyncJobDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAsyncJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncJobDetailRequest", string(data)}, " ")
}

type ListAsyncJobDetailRequestXLanguage struct {
	value string
}

type ListAsyncJobDetailRequestXLanguageEnum struct {
	EN_US ListAsyncJobDetailRequestXLanguage
	ZH_CN ListAsyncJobDetailRequestXLanguage
}

func GetListAsyncJobDetailRequestXLanguageEnum() ListAsyncJobDetailRequestXLanguageEnum {
	return ListAsyncJobDetailRequestXLanguageEnum{
		EN_US: ListAsyncJobDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAsyncJobDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAsyncJobDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListAsyncJobDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAsyncJobDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
