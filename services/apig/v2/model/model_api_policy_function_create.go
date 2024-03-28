package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyFunctionCreate struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiPolicyFunctionCreateInvocationType `json:"invocation_type"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType ApiPolicyFunctionCreateNetworkType `json:"network_type"`

	// 函数版本  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// API网关请求后端服务的超时时间。函数网络架构为V1时最大超时时间为60000，V2最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 函数后端的请求协议：HTTPS、GRPCS，默认值为HTTPS，前端配置中的请求协议为GRPCS时可选GRPCS。
	ReqProtocol *ApiPolicyFunctionCreateReqProtocol `json:"req_protocol,omitempty"`

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyFunctionCreateEffectMode `json:"effect_mode"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name"`

	// 后端参数列表，后端类型为GRPC时不支持配置
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`

	// 策略条件列表
	Conditions []ApiConditionBase `json:"conditions"`

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

type ApiPolicyFunctionCreateNetworkType struct {
	value string
}

type ApiPolicyFunctionCreateNetworkTypeEnum struct {
	V1 ApiPolicyFunctionCreateNetworkType
	V2 ApiPolicyFunctionCreateNetworkType
}

func GetApiPolicyFunctionCreateNetworkTypeEnum() ApiPolicyFunctionCreateNetworkTypeEnum {
	return ApiPolicyFunctionCreateNetworkTypeEnum{
		V1: ApiPolicyFunctionCreateNetworkType{
			value: "V1",
		},
		V2: ApiPolicyFunctionCreateNetworkType{
			value: "V2",
		},
	}
}

func (c ApiPolicyFunctionCreateNetworkType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionCreateNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionCreateNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionCreateReqProtocol struct {
	value string
}

type ApiPolicyFunctionCreateReqProtocolEnum struct {
	HTTPS ApiPolicyFunctionCreateReqProtocol
	GRPCS ApiPolicyFunctionCreateReqProtocol
}

func GetApiPolicyFunctionCreateReqProtocolEnum() ApiPolicyFunctionCreateReqProtocolEnum {
	return ApiPolicyFunctionCreateReqProtocolEnum{
		HTTPS: ApiPolicyFunctionCreateReqProtocol{
			value: "HTTPS",
		},
		GRPCS: ApiPolicyFunctionCreateReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiPolicyFunctionCreateReqProtocol) Value() string {
	return c.value
}

func (c ApiPolicyFunctionCreateReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionCreateReqProtocol) UnmarshalJSON(b []byte) error {
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
