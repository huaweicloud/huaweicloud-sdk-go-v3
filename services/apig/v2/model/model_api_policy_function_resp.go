package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyFunctionResp struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiPolicyFunctionRespInvocationType `json:"invocation_type"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType ApiPolicyFunctionRespNetworkType `json:"network_type"`

	// 函数版本  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// API网关请求后端服务的超时时间。函数网络架构为V1时最大超时时间为60000，V2最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 函数后端的请求协议：HTTPS、GRPCS，默认值为HTTPS，前端配置中的请求协议为GRPCS时可选GRPCS。
	ReqProtocol *ApiPolicyFunctionRespReqProtocol `json:"req_protocol,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyFunctionRespEffectMode `json:"effect_mode"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name"`

	// 后端参数列表
	BackendParams *[]BackendParam `json:"backend_params,omitempty"`

	// 策略条件列表
	Conditions []CoditionResp `json:"conditions"`

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyFunctionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyFunctionResp struct{}"
	}

	return strings.Join([]string{"ApiPolicyFunctionResp", string(data)}, " ")
}

type ApiPolicyFunctionRespInvocationType struct {
	value string
}

type ApiPolicyFunctionRespInvocationTypeEnum struct {
	ASYNC ApiPolicyFunctionRespInvocationType
	SYNC  ApiPolicyFunctionRespInvocationType
}

func GetApiPolicyFunctionRespInvocationTypeEnum() ApiPolicyFunctionRespInvocationTypeEnum {
	return ApiPolicyFunctionRespInvocationTypeEnum{
		ASYNC: ApiPolicyFunctionRespInvocationType{
			value: "async",
		},
		SYNC: ApiPolicyFunctionRespInvocationType{
			value: "sync",
		},
	}
}

func (c ApiPolicyFunctionRespInvocationType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionRespNetworkType struct {
	value string
}

type ApiPolicyFunctionRespNetworkTypeEnum struct {
	V1 ApiPolicyFunctionRespNetworkType
	V2 ApiPolicyFunctionRespNetworkType
}

func GetApiPolicyFunctionRespNetworkTypeEnum() ApiPolicyFunctionRespNetworkTypeEnum {
	return ApiPolicyFunctionRespNetworkTypeEnum{
		V1: ApiPolicyFunctionRespNetworkType{
			value: "V1",
		},
		V2: ApiPolicyFunctionRespNetworkType{
			value: "V2",
		},
	}
}

func (c ApiPolicyFunctionRespNetworkType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionRespReqProtocol struct {
	value string
}

type ApiPolicyFunctionRespReqProtocolEnum struct {
	HTTPS ApiPolicyFunctionRespReqProtocol
	GRPCS ApiPolicyFunctionRespReqProtocol
}

func GetApiPolicyFunctionRespReqProtocolEnum() ApiPolicyFunctionRespReqProtocolEnum {
	return ApiPolicyFunctionRespReqProtocolEnum{
		HTTPS: ApiPolicyFunctionRespReqProtocol{
			value: "HTTPS",
		},
		GRPCS: ApiPolicyFunctionRespReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiPolicyFunctionRespReqProtocol) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionRespEffectMode struct {
	value string
}

type ApiPolicyFunctionRespEffectModeEnum struct {
	ALL ApiPolicyFunctionRespEffectMode
	ANY ApiPolicyFunctionRespEffectMode
}

func GetApiPolicyFunctionRespEffectModeEnum() ApiPolicyFunctionRespEffectModeEnum {
	return ApiPolicyFunctionRespEffectModeEnum{
		ALL: ApiPolicyFunctionRespEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyFunctionRespEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyFunctionRespEffectMode) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespEffectMode) UnmarshalJSON(b []byte) error {
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
