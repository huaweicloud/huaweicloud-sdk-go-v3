package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyHttpBase struct {

	// 策略后端的Endpoint。  由域名（或IP地址）和端口号组成，总长度不超过255。格式为域名:端口（如：apig.example.com:7443）。如果不写端口，则HTTPS默认端口号为443， HTTP默认端口号为80。  支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、“_”、“-”组成，且只能以英文开头。
	UrlDomain *string `json:"url_domain,omitempty"`

	// 请求协议：HTTP、HTTPS、GRPC、GRPCS，后端类型为GRPC&GRPCS时可选GRPC、GRPCS，当vpc_channel_status取值为1或者2时，该字段必填。
	ReqProtocol *ApiPolicyHttpBaseReqProtocol `json:"req_protocol,omitempty"`

	// 请求方式：GET、POST、PUT、DELETE、HEAD、PATCH、OPTIONS、ANY，后端类型为GRPC&GRPCS时固定为POST，当vpc_channel_status取值为1或者2时，该字段必填。
	ReqMethod *ApiPolicyHttpBaseReqMethod `json:"req_method,omitempty"`

	// 请求地址，当vpc_channel_status取值为1或者2时，该字段必填。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512字符，且满足URI规范。   支持环境变量，使用环境变量时，每个变量名的长度为3~32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。  > 需要服从URI规范。  后端类型为GRPC时请求地址固定为/
	ReqUri *string `json:"req_uri,omitempty"`

	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 请求后端服务的重试次数，默认为-1，范围[-1,10]。  当该值为-1时，幂等的接口会重试1次，非幂等的不会重试。POST，PATCH方法为非幂等；GET，HEAD，PUT，OPTIONS和DELETE等方法为幂等的。
	RetryCount *string `json:"retry_count,omitempty"`

	// 是否启用SM商密通道。  仅实例支持SM系列商密算法的实例时支持开启。
	EnableSmChannel *bool `json:"enable_sm_channel,omitempty"`

	// 后端服务器分组详细信息，当vpc_channel_status取值为4时，该字段必填。
	MemberGroupUrlInfos *[]MemberGroupUrlInfo `json:"member_group_url_infos,omitempty"`
}

func (o ApiPolicyHttpBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyHttpBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyHttpBase", string(data)}, " ")
}

type ApiPolicyHttpBaseReqProtocol struct {
	value string
}

type ApiPolicyHttpBaseReqProtocolEnum struct {
	HTTP  ApiPolicyHttpBaseReqProtocol
	HTTPS ApiPolicyHttpBaseReqProtocol
	GRPC  ApiPolicyHttpBaseReqProtocol
	GRPCS ApiPolicyHttpBaseReqProtocol
}

func GetApiPolicyHttpBaseReqProtocolEnum() ApiPolicyHttpBaseReqProtocolEnum {
	return ApiPolicyHttpBaseReqProtocolEnum{
		HTTP: ApiPolicyHttpBaseReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiPolicyHttpBaseReqProtocol{
			value: "HTTPS",
		},
		GRPC: ApiPolicyHttpBaseReqProtocol{
			value: "GRPC",
		},
		GRPCS: ApiPolicyHttpBaseReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiPolicyHttpBaseReqProtocol) Value() string {
	return c.value
}

func (c ApiPolicyHttpBaseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyHttpBaseReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiPolicyHttpBaseReqMethod struct {
	value string
}

type ApiPolicyHttpBaseReqMethodEnum struct {
	GET     ApiPolicyHttpBaseReqMethod
	POST    ApiPolicyHttpBaseReqMethod
	PUT     ApiPolicyHttpBaseReqMethod
	DELETE  ApiPolicyHttpBaseReqMethod
	HEAD    ApiPolicyHttpBaseReqMethod
	PATCH   ApiPolicyHttpBaseReqMethod
	OPTIONS ApiPolicyHttpBaseReqMethod
	ANY     ApiPolicyHttpBaseReqMethod
}

func GetApiPolicyHttpBaseReqMethodEnum() ApiPolicyHttpBaseReqMethodEnum {
	return ApiPolicyHttpBaseReqMethodEnum{
		GET: ApiPolicyHttpBaseReqMethod{
			value: "GET",
		},
		POST: ApiPolicyHttpBaseReqMethod{
			value: "POST",
		},
		PUT: ApiPolicyHttpBaseReqMethod{
			value: "PUT",
		},
		DELETE: ApiPolicyHttpBaseReqMethod{
			value: "DELETE",
		},
		HEAD: ApiPolicyHttpBaseReqMethod{
			value: "HEAD",
		},
		PATCH: ApiPolicyHttpBaseReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiPolicyHttpBaseReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiPolicyHttpBaseReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiPolicyHttpBaseReqMethod) Value() string {
	return c.value
}

func (c ApiPolicyHttpBaseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyHttpBaseReqMethod) UnmarshalJSON(b []byte) error {
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
