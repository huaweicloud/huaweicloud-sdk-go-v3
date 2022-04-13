package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// web后端详情
type BackendApiCreate struct {
	// 后端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 后端服务的地址。  由主机（IP或域名）和端口号组成，总长度不超过255。格式为主机:端口（如：apig.example.com:7443）。如果不写端口，则HTTPS默认端口号为443，HTTP默认端口号为80。  支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、下划线、中划线组成，且只能以英文开头

	UrlDomain *string `json:"url_domain,omitempty"`
	// 请求协议

	ReqProtocol BackendApiCreateReqProtocol `json:"req_protocol"`
	// 描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// 请求方式

	ReqMethod BackendApiCreateReqMethod `json:"req_method"`
	// web后端版本，字符长度不超过16

	Version *string `json:"version,omitempty"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。

	Timeout int32 `json:"timeout"`
	// 是否开启双向认证

	EnableClientSsl *bool `json:"enable_client_ssl,omitempty"`

	VpcChannelInfo *ApiBackendVpcReq `json:"vpc_channel_info,omitempty"`
	// 是否使用VPC通道 - 1：使用VPC通道 - 2：不使用VPC通道

	VpcChannelStatus *BackendApiCreateVpcChannelStatus `json:"vpc_channel_status,omitempty"`
}

func (o BackendApiCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendApiCreate struct{}"
	}

	return strings.Join([]string{"BackendApiCreate", string(data)}, " ")
}

type BackendApiCreateReqProtocol struct {
	value string
}

type BackendApiCreateReqProtocolEnum struct {
	HTTP  BackendApiCreateReqProtocol
	HTTPS BackendApiCreateReqProtocol
}

func GetBackendApiCreateReqProtocolEnum() BackendApiCreateReqProtocolEnum {
	return BackendApiCreateReqProtocolEnum{
		HTTP: BackendApiCreateReqProtocol{
			value: "HTTP",
		},
		HTTPS: BackendApiCreateReqProtocol{
			value: "HTTPS",
		},
	}
}

func (c BackendApiCreateReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiCreateReqProtocol) UnmarshalJSON(b []byte) error {
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

type BackendApiCreateReqMethod struct {
	value string
}

type BackendApiCreateReqMethodEnum struct {
	GET     BackendApiCreateReqMethod
	POST    BackendApiCreateReqMethod
	PUT     BackendApiCreateReqMethod
	DELETE  BackendApiCreateReqMethod
	HEAD    BackendApiCreateReqMethod
	PATCH   BackendApiCreateReqMethod
	OPTIONS BackendApiCreateReqMethod
	ANY     BackendApiCreateReqMethod
}

func GetBackendApiCreateReqMethodEnum() BackendApiCreateReqMethodEnum {
	return BackendApiCreateReqMethodEnum{
		GET: BackendApiCreateReqMethod{
			value: "GET",
		},
		POST: BackendApiCreateReqMethod{
			value: "POST",
		},
		PUT: BackendApiCreateReqMethod{
			value: "PUT",
		},
		DELETE: BackendApiCreateReqMethod{
			value: "DELETE",
		},
		HEAD: BackendApiCreateReqMethod{
			value: "HEAD",
		},
		PATCH: BackendApiCreateReqMethod{
			value: "PATCH",
		},
		OPTIONS: BackendApiCreateReqMethod{
			value: "OPTIONS",
		},
		ANY: BackendApiCreateReqMethod{
			value: "ANY",
		},
	}
}

func (c BackendApiCreateReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiCreateReqMethod) UnmarshalJSON(b []byte) error {
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

type BackendApiCreateVpcChannelStatus struct {
	value int32
}

type BackendApiCreateVpcChannelStatusEnum struct {
	E_1 BackendApiCreateVpcChannelStatus
	E_2 BackendApiCreateVpcChannelStatus
}

func GetBackendApiCreateVpcChannelStatusEnum() BackendApiCreateVpcChannelStatusEnum {
	return BackendApiCreateVpcChannelStatusEnum{
		E_1: BackendApiCreateVpcChannelStatus{
			value: 1,
		}, E_2: BackendApiCreateVpcChannelStatus{
			value: 2,
		},
	}
}

func (c BackendApiCreateVpcChannelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendApiCreateVpcChannelStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
