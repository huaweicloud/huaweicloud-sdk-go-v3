package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ApiFunc 函数工作流后端详情
type ApiFunc struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiFuncInvocationType `json:"invocation_type"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType ApiFuncNetworkType `json:"network_type"`

	// 函数版本  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// API网关请求后端服务的超时时间。函数网络架构为V1时最大超时时间为60000，V2最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 函数后端的请求协议：HTTPS、GRPCS，默认值为HTTPS，前端配置中的请求协议为GRPCS时可选GRPCS。
	ReqProtocol *ApiFuncReqProtocol `json:"req_protocol,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 后端状态   - 1： 有效
	Status *int32 `json:"status,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiFunc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiFunc struct{}"
	}

	return strings.Join([]string{"ApiFunc", string(data)}, " ")
}

type ApiFuncInvocationType struct {
	value string
}

type ApiFuncInvocationTypeEnum struct {
	ASYNC ApiFuncInvocationType
	SYNC  ApiFuncInvocationType
}

func GetApiFuncInvocationTypeEnum() ApiFuncInvocationTypeEnum {
	return ApiFuncInvocationTypeEnum{
		ASYNC: ApiFuncInvocationType{
			value: "async",
		},
		SYNC: ApiFuncInvocationType{
			value: "sync",
		},
	}
}

func (c ApiFuncInvocationType) Value() string {
	return c.value
}

func (c ApiFuncInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiFuncNetworkType struct {
	value string
}

type ApiFuncNetworkTypeEnum struct {
	V1 ApiFuncNetworkType
	V2 ApiFuncNetworkType
}

func GetApiFuncNetworkTypeEnum() ApiFuncNetworkTypeEnum {
	return ApiFuncNetworkTypeEnum{
		V1: ApiFuncNetworkType{
			value: "V1",
		},
		V2: ApiFuncNetworkType{
			value: "V2",
		},
	}
}

func (c ApiFuncNetworkType) Value() string {
	return c.value
}

func (c ApiFuncNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiFuncReqProtocol struct {
	value string
}

type ApiFuncReqProtocolEnum struct {
	HTTPS ApiFuncReqProtocol
	GRPCS ApiFuncReqProtocol
}

func GetApiFuncReqProtocolEnum() ApiFuncReqProtocolEnum {
	return ApiFuncReqProtocolEnum{
		HTTPS: ApiFuncReqProtocol{
			value: "HTTPS",
		},
		GRPCS: ApiFuncReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiFuncReqProtocol) Value() string {
	return c.value
}

func (c ApiFuncReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncReqProtocol) UnmarshalJSON(b []byte) error {
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
