package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyFunctionCreate struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiPolicyFunctionCreateInvocationType `json:"invocation_type"`

	// ROMA Connect APIC请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyFunctionCreateEffectMode `json:"effect_mode"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name"`

	// 后端参数列表
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`

	// 策略条件列表
	Conditions []ApiConditionCreate `json:"conditions"`

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyFunctionCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyFunctionCreate struct{}"
	}

	return strings.Join([]string{"ApiPolicyFunctionCreate", string(data)}, " ")
}

type ApiPolicyFunctionCreateInvocationType struct {
	value string
}

type ApiPolicyFunctionCreateInvocationTypeEnum struct {
	ASYNC ApiPolicyFunctionCreateInvocationType
	SYNC  ApiPolicyFunctionCreateInvocationType
}

func GetApiPolicyFunctionCreateInvocationTypeEnum() ApiPolicyFunctionCreateInvocationTypeEnum {
	return ApiPolicyFunctionCreateInvocationTypeEnum{
		ASYNC: ApiPolicyFunctionCreateInvocationType{
			value: "async",
		},
		SYNC: ApiPolicyFunctionCreateInvocationType{
			value: "sync",
		},
	}
}

func (c ApiPolicyFunctionCreateInvocationType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionCreateInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionCreateInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionCreateEffectMode struct {
	value string
}

type ApiPolicyFunctionCreateEffectModeEnum struct {
	ALL ApiPolicyFunctionCreateEffectMode
	ANY ApiPolicyFunctionCreateEffectMode
}

func GetApiPolicyFunctionCreateEffectModeEnum() ApiPolicyFunctionCreateEffectModeEnum {
	return ApiPolicyFunctionCreateEffectModeEnum{
		ALL: ApiPolicyFunctionCreateEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyFunctionCreateEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyFunctionCreateEffectMode) Value() string {
	return c.value
}

func (c ApiPolicyFunctionCreateEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionCreateEffectMode) UnmarshalJSON(b []byte) error {
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
