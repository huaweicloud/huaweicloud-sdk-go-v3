package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApiFuncCreate 函数后端详情
type ApiFuncCreate struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiFuncCreateInvocationType `json:"invocation_type"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType ApiFuncCreateNetworkType `json:"network_type"`

	// 函数版本  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// API网关请求后端服务的超时时间。函数网络架构为V1时最大超时时间为60000，V2最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 函数后端的请求协议：HTTPS、GRPCS，默认值为HTTPS，前端配置中的请求协议为GRPCS时可选GRPCS。
	ReqProtocol *ApiFuncCreateReqProtocol `json:"req_protocol,omitempty"`
}

func (o ApiFuncCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiFuncCreate struct{}"
	}

	return strings.Join([]string{"ApiFuncCreate", string(data)}, " ")
}

type ApiFuncCreateInvocationType struct {
	value string
}

type ApiFuncCreateInvocationTypeEnum struct {
	ASYNC ApiFuncCreateInvocationType
	SYNC  ApiFuncCreateInvocationType
}

func GetApiFuncCreateInvocationTypeEnum() ApiFuncCreateInvocationTypeEnum {
	return ApiFuncCreateInvocationTypeEnum{
		ASYNC: ApiFuncCreateInvocationType{
			value: "async",
		},
		SYNC: ApiFuncCreateInvocationType{
			value: "sync",
		},
	}
}

func (c ApiFuncCreateInvocationType) Value() string {
	return c.value
}

func (c ApiFuncCreateInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncCreateInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiFuncCreateNetworkType struct {
	value string
}

type ApiFuncCreateNetworkTypeEnum struct {
	V1 ApiFuncCreateNetworkType
	V2 ApiFuncCreateNetworkType
}

func GetApiFuncCreateNetworkTypeEnum() ApiFuncCreateNetworkTypeEnum {
	return ApiFuncCreateNetworkTypeEnum{
		V1: ApiFuncCreateNetworkType{
			value: "V1",
		},
		V2: ApiFuncCreateNetworkType{
			value: "V2",
		},
	}
}

func (c ApiFuncCreateNetworkType) Value() string {
	return c.value
}

func (c ApiFuncCreateNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncCreateNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiFuncCreateReqProtocol struct {
	value string
}

type ApiFuncCreateReqProtocolEnum struct {
	HTTPS ApiFuncCreateReqProtocol
	GRPCS ApiFuncCreateReqProtocol
}

func GetApiFuncCreateReqProtocolEnum() ApiFuncCreateReqProtocolEnum {
	return ApiFuncCreateReqProtocolEnum{
		HTTPS: ApiFuncCreateReqProtocol{
			value: "HTTPS",
		},
		GRPCS: ApiFuncCreateReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiFuncCreateReqProtocol) Value() string {
	return c.value
}

func (c ApiFuncCreateReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncCreateReqProtocol) UnmarshalJSON(b []byte) error {
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
