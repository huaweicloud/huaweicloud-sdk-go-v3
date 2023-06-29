package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchExecuteJobActionsRequest Request Object
type BatchExecuteJobActionsRequest struct {

	// 请求语言类型。
	XLanguage *BatchExecuteJobActionsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchJobActionReq `json:"body,omitempty"`
}

func (o BatchExecuteJobActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteJobActionsRequest struct{}"
	}

	return strings.Join([]string{"BatchExecuteJobActionsRequest", string(data)}, " ")
}

type BatchExecuteJobActionsRequestXLanguage struct {
	value string
}

type BatchExecuteJobActionsRequestXLanguageEnum struct {
	EN_US BatchExecuteJobActionsRequestXLanguage
	ZH_CN BatchExecuteJobActionsRequestXLanguage
}

func GetBatchExecuteJobActionsRequestXLanguageEnum() BatchExecuteJobActionsRequestXLanguageEnum {
	return BatchExecuteJobActionsRequestXLanguageEnum{
		EN_US: BatchExecuteJobActionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchExecuteJobActionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchExecuteJobActionsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchExecuteJobActionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchExecuteJobActionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
