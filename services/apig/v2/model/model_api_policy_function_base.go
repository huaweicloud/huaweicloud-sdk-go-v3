package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyFunctionBase struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiPolicyFunctionBaseInvocationType `json:"invocation_type"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o ApiPolicyFunctionBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyFunctionBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyFunctionBase", string(data)}, " ")
}

type ApiPolicyFunctionBaseInvocationType struct {
	value string
}

type ApiPolicyFunctionBaseInvocationTypeEnum struct {
	ASYNC ApiPolicyFunctionBaseInvocationType
	SYNC  ApiPolicyFunctionBaseInvocationType
}

func GetApiPolicyFunctionBaseInvocationTypeEnum() ApiPolicyFunctionBaseInvocationTypeEnum {
	return ApiPolicyFunctionBaseInvocationTypeEnum{
		ASYNC: ApiPolicyFunctionBaseInvocationType{
			value: "async",
		},
		SYNC: ApiPolicyFunctionBaseInvocationType{
			value: "sync",
		},
	}
}

func (c ApiPolicyFunctionBaseInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionBaseInvocationType) UnmarshalJSON(b []byte) error {
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
