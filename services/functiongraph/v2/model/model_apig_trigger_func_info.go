package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApigTriggerFuncInfo APIG触发器函数工作流后端详情（APIG触发器参数）。APIG触发器此参数必填。
type ApigTriggerFuncInfo struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn *string `json:"function_urn,omitempty"`

	// 调用函数执行方式。 - sync：同步执行 - async：异步执行
	InvocationType *ApigTriggerFuncInfoInvocationType `json:"invocation_type,omitempty"`

	// API网关请求函数服务的超时时间（毫秒）。APIG触发器此参数必填。
	Timeout int32 `json:"timeout"`

	// 函数版本信息。
	Version *string `json:"version,omitempty"`
}

func (o ApigTriggerFuncInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigTriggerFuncInfo struct{}"
	}

	return strings.Join([]string{"ApigTriggerFuncInfo", string(data)}, " ")
}

type ApigTriggerFuncInfoInvocationType struct {
	value string
}

type ApigTriggerFuncInfoInvocationTypeEnum struct {
	SYNC  ApigTriggerFuncInfoInvocationType
	ASYNC ApigTriggerFuncInfoInvocationType
}

func GetApigTriggerFuncInfoInvocationTypeEnum() ApigTriggerFuncInfoInvocationTypeEnum {
	return ApigTriggerFuncInfoInvocationTypeEnum{
		SYNC: ApigTriggerFuncInfoInvocationType{
			value: "sync",
		},
		ASYNC: ApigTriggerFuncInfoInvocationType{
			value: "async",
		},
	}
}

func (c ApigTriggerFuncInfoInvocationType) Value() string {
	return c.value
}

func (c ApigTriggerFuncInfoInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApigTriggerFuncInfoInvocationType) UnmarshalJSON(b []byte) error {
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
