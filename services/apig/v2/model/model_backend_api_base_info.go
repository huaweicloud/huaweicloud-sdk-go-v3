package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendApiBaseInfo struct {

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 后端服务的地址。   由主机（IP或域名）和端口号组成，总长度不超过255。格式为主机:端口（如：apig.example.com:7443）。如果不写端口，则HTTPS默认端口号为443，HTTP默认端口号为80。   支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、下划线、中划线组成，且只能以英文开头
	UrlDomain *string `json:"url_domain,omitempty"`

	// 请求协议
	ReqProtocol BackendApiBaseInfoReqProtocol `json:"req_protocol"`

	// 描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 请求方式
	ReqMethod BackendApiBaseInfoReqMethod `json:"req_method"`

	// web后端版本，字符长度不超过16
	Version *string `json:"version,omitempty"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。   支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 是否开启双向认证
	EnableClientSsl *bool `json:"enable_client_ssl,omitempty"`

	// 请求后端服务的重试次数，默认为-1，范围[-1,10]。  当该值为-1时，幂等的接口会重试1次，非幂等的不会重试。POST，PATCH方法为非幂等；GET，HEAD，PUT，OPTIONS和DELETE等方法为幂等的。
	RetryCount *string `json:"retry_count,omitempty"`
}

func (o BackendApiBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendApiBaseInfo struct{}"
	}

	return strings.Join([]string{"BackendApiBaseInfo", string(data)}, " ")
}

type BackendApiBaseInfoReqProtocol struct {
	value string
}

type BackendApiBaseInfoReqProtocolEnum struct {
	HTTP  BackendApiBaseInfoReqProtocol
	HTTPS BackendApiBaseInfoReqProtocol
}

func GetBackendApiBaseInfoReqProtocolEnum() BackendApiBaseInfoReqProtocolEnum {
	return BackendApiBaseInfoReqProtocolEnum{
		HTTP: BackendApiBaseInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: BackendApiBaseInfoReqProtocol{
			value: "HTTPS",
		},
	}
}

func (c BackendApiBaseInfoReqProtocol) Value() string {
	return c.value
}

func (c BackendApiBaseInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiBaseInfoReqProtocol) UnmarshalJSON(b []byte) error {
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

type BackendApiBaseInfoReqMethod struct {
	value string
}

type BackendApiBaseInfoReqMethodEnum struct {
	GET     BackendApiBaseInfoReqMethod
	POST    BackendApiBaseInfoReqMethod
	PUT     BackendApiBaseInfoReqMethod
	DELETE  BackendApiBaseInfoReqMethod
	HEAD    BackendApiBaseInfoReqMethod
	PATCH   BackendApiBaseInfoReqMethod
	OPTIONS BackendApiBaseInfoReqMethod
	ANY     BackendApiBaseInfoReqMethod
}

func GetBackendApiBaseInfoReqMethodEnum() BackendApiBaseInfoReqMethodEnum {
	return BackendApiBaseInfoReqMethodEnum{
		GET: BackendApiBaseInfoReqMethod{
			value: "GET",
		},
		POST: BackendApiBaseInfoReqMethod{
			value: "POST",
		},
		PUT: BackendApiBaseInfoReqMethod{
			value: "PUT",
		},
		DELETE: BackendApiBaseInfoReqMethod{
			value: "DELETE",
		},
		HEAD: BackendApiBaseInfoReqMethod{
			value: "HEAD",
		},
		PATCH: BackendApiBaseInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: BackendApiBaseInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: BackendApiBaseInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c BackendApiBaseInfoReqMethod) Value() string {
	return c.value
}

func (c BackendApiBaseInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiBaseInfoReqMethod) UnmarshalJSON(b []byte) error {
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
