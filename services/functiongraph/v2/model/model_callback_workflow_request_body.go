package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CallbackWorkflowRequestBody 回调工作流的请求Body
type CallbackWorkflowRequestBody struct {

	// 执行结果
	Result CallbackWorkflowRequestBodyResult `json:"result"`

	// 错误信息
	Error *string `json:"error,omitempty"`

	// 工作流的执行结果，JSON格式，仅在status为success时有值
	Output *interface{} `json:"output"`
}

func (o CallbackWorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CallbackWorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"CallbackWorkflowRequestBody", string(data)}, " ")
}

type CallbackWorkflowRequestBodyResult struct {
	value string
}

type CallbackWorkflowRequestBodyResultEnum struct {
	SUCCESS CallbackWorkflowRequestBodyResult
	FAIL    CallbackWorkflowRequestBodyResult
}

func GetCallbackWorkflowRequestBodyResultEnum() CallbackWorkflowRequestBodyResultEnum {
	return CallbackWorkflowRequestBodyResultEnum{
		SUCCESS: CallbackWorkflowRequestBodyResult{
			value: "success",
		},
		FAIL: CallbackWorkflowRequestBodyResult{
			value: "fail",
		},
	}
}

func (c CallbackWorkflowRequestBodyResult) Value() string {
	return c.value
}

func (c CallbackWorkflowRequestBodyResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CallbackWorkflowRequestBodyResult) UnmarshalJSON(b []byte) error {
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
