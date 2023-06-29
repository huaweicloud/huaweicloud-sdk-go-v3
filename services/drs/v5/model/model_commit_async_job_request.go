package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommitAsyncJobRequest Request Object
type CommitAsyncJobRequest struct {

	// 批量异步创建的任务ID，由创建批量异步任务接口返回。
	AsyncJobId string `json:"async_job_id"`

	// 请求语言类型。
	XLanguage *CommitAsyncJobRequestXLanguage `json:"X-Language,omitempty"`
}

func (o CommitAsyncJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitAsyncJobRequest struct{}"
	}

	return strings.Join([]string{"CommitAsyncJobRequest", string(data)}, " ")
}

type CommitAsyncJobRequestXLanguage struct {
	value string
}

type CommitAsyncJobRequestXLanguageEnum struct {
	EN_US CommitAsyncJobRequestXLanguage
	ZH_CN CommitAsyncJobRequestXLanguage
}

func GetCommitAsyncJobRequestXLanguageEnum() CommitAsyncJobRequestXLanguageEnum {
	return CommitAsyncJobRequestXLanguageEnum{
		EN_US: CommitAsyncJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CommitAsyncJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CommitAsyncJobRequestXLanguage) Value() string {
	return c.value
}

func (c CommitAsyncJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommitAsyncJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
