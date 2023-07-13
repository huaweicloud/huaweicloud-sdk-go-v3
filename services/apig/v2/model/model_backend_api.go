package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// BackendApi web后端详情
type BackendApi struct {

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 后端服务的地址。   由主机（IP或域名）和端口号组成，总长度不超过255。格式为主机:端口（如：apig.example.com:7443）。如果不写端口，则HTTPS默认端口号为443，HTTP默认端口号为80。   支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、下划线、中划线组成，且只能以英文开头
	UrlDomain *string `json:"url_domain,omitempty"`

	// 请求协议
	ReqProtocol BackendApiReqProtocol `json:"req_protocol"`

	// 描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 请求方式
	ReqMethod BackendApiReqMethod `json:"req_method"`

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

	// 编号
	Id *string `json:"id,omitempty"`

	// 后端状态   - 1： 有效
	Status *int32 `json:"status,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	VpcChannelInfo *VpcInfo `json:"vpc_channel_info,omitempty"`

	// 是否使用VPC通道 - 1：使用VPC通道 - 2：不使用VPC通道
	VpcChannelStatus *int32 `json:"vpc_channel_status,omitempty"`
}

func (o BackendApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendApi struct{}"
	}

	return strings.Join([]string{"BackendApi", string(data)}, " ")
}

type BackendApiReqProtocol struct {
	value string
}

type BackendApiReqProtocolEnum struct {
	HTTP  BackendApiReqProtocol
	HTTPS BackendApiReqProtocol
}

func GetBackendApiReqProtocolEnum() BackendApiReqProtocolEnum {
	return BackendApiReqProtocolEnum{
		HTTP: BackendApiReqProtocol{
			value: "HTTP",
		},
		HTTPS: BackendApiReqProtocol{
			value: "HTTPS",
		},
	}
}

func (c BackendApiReqProtocol) Value() string {
	return c.value
}

func (c BackendApiReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiReqProtocol) UnmarshalJSON(b []byte) error {
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

type BackendApiReqMethod struct {
	value string
}

type BackendApiReqMethodEnum struct {
	GET     BackendApiReqMethod
	POST    BackendApiReqMethod
	PUT     BackendApiReqMethod
	DELETE  BackendApiReqMethod
	HEAD    BackendApiReqMethod
	PATCH   BackendApiReqMethod
	OPTIONS BackendApiReqMethod
	ANY     BackendApiReqMethod
}

func GetBackendApiReqMethodEnum() BackendApiReqMethodEnum {
	return BackendApiReqMethodEnum{
		GET: BackendApiReqMethod{
			value: "GET",
		},
		POST: BackendApiReqMethod{
			value: "POST",
		},
		PUT: BackendApiReqMethod{
			value: "PUT",
		},
		DELETE: BackendApiReqMethod{
			value: "DELETE",
		},
		HEAD: BackendApiReqMethod{
			value: "HEAD",
		},
		PATCH: BackendApiReqMethod{
			value: "PATCH",
		},
		OPTIONS: BackendApiReqMethod{
			value: "OPTIONS",
		},
		ANY: BackendApiReqMethod{
			value: "ANY",
		},
	}
}

func (c BackendApiReqMethod) Value() string {
	return c.value
}

func (c BackendApiReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiReqMethod) UnmarshalJSON(b []byte) error {
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
