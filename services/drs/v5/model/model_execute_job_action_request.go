package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteJobActionRequest Request Object
type ExecuteJobActionRequest struct {

	// 任务ID (对比任务相关操作，多任务场景传父任务详情返回的master_job_id)
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ExecuteJobActionRequestXLanguage `json:"X-Language,omitempty"`

	Body *JobActionReq `json:"body,omitempty"`
}

func (o ExecuteJobActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteJobActionRequest", string(data)}, " ")
}

type ExecuteJobActionRequestXLanguage struct {
	value string
}

type ExecuteJobActionRequestXLanguageEnum struct {
	EN_US ExecuteJobActionRequestXLanguage
	ZH_CN ExecuteJobActionRequestXLanguage
}

func GetExecuteJobActionRequestXLanguageEnum() ExecuteJobActionRequestXLanguageEnum {
	return ExecuteJobActionRequestXLanguageEnum{
		EN_US: ExecuteJobActionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExecuteJobActionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExecuteJobActionRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteJobActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteJobActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
