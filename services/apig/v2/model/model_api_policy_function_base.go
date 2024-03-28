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

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType ApiPolicyFunctionBaseNetworkType `json:"network_type"`

	// 函数版本  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// API网关请求后端服务的超时时间。函数网络架构为V1时最大超时时间为60000，V2最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 函数后端的请求协议：HTTPS、GRPCS，默认值为HTTPS，前端配置中的请求协议为GRPCS时可选GRPCS。
	ReqProtocol *ApiPolicyFunctionBaseReqProtocol `json:"req_protocol,omitempty"`
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

func (c ApiPolicyFunctionBaseInvocationType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionBaseInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionBaseInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionBaseNetworkType struct {
	value string
}

type ApiPolicyFunctionBaseNetworkTypeEnum struct {
	V1 ApiPolicyFunctionBaseNetworkType
	V2 ApiPolicyFunctionBaseNetworkType
}

func GetApiPolicyFunctionBaseNetworkTypeEnum() ApiPolicyFunctionBaseNetworkTypeEnum {
	return ApiPolicyFunctionBaseNetworkTypeEnum{
		V1: ApiPolicyFunctionBaseNetworkType{
			value: "V1",
		},
		V2: ApiPolicyFunctionBaseNetworkType{
			value: "V2",
		},
	}
}

func (c ApiPolicyFunctionBaseNetworkType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionBaseNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionBaseNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionBaseReqProtocol struct {
	value string
}

type ApiPolicyFunctionBaseReqProtocolEnum struct {
	HTTPS ApiPolicyFunctionBaseReqProtocol
	GRPCS ApiPolicyFunctionBaseReqProtocol
}

func GetApiPolicyFunctionBaseReqProtocolEnum() ApiPolicyFunctionBaseReqProtocolEnum {
	return ApiPolicyFunctionBaseReqProtocolEnum{
		HTTPS: ApiPolicyFunctionBaseReqProtocol{
			value: "HTTPS",
		},
		GRPCS: ApiPolicyFunctionBaseReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiPolicyFunctionBaseReqProtocol) Value() string {
	return c.value
}

func (c ApiPolicyFunctionBaseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionBaseReqProtocol) UnmarshalJSON(b []byte) error {
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
