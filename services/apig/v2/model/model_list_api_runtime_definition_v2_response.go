package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ListApiRuntimeDefinitionV2Response struct {
	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型 - 1：公有API - 2：私有API

	Type ListApiRuntimeDefinitionV2ResponseType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol ListApiRuntimeDefinitionV2ResponseReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod ListApiRuntimeDefinitionV2ResponseReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType ListApiRuntimeDefinitionV2ResponseAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *ListApiRuntimeDefinitionV2ResponseMatchMode `json:"match_mode,omitempty"`
	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType ListApiRuntimeDefinitionV2ResponseBackendType `json:"backend_type"`
	// API描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// API所属的分组编号

	GroupId string `json:"group_id"`
	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。

	BodyRemark *string `json:"body_remark,omitempty"`
	// 正常响应示例，描述API的正常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。

	ResultNormalSample *string `json:"result_normal_sample,omitempty"`
	// 失败返回示例，描述API的异常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。

	ResultFailureSample *string `json:"result_failure_sample,omitempty"`
	// 前端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 标签。  支持英文，数字，下划线，且只能以英文开头。支持输入多个标签，不同标签以英文逗号分割。

	Tags *[]string `json:"tags,omitempty"`
	// 分组自定义响应ID

	ResponseId *string `json:"response_id,omitempty"`
	// 集成应用ID  暂不支持

	RomaAppId *string `json:"roma_app_id,omitempty"`
	// API绑定的自定义域名  暂不支持

	DomainName *string `json:"domain_name,omitempty"`
	// 标签  待废弃，优先使用tags字段

	Tag *string `json:"tag,omitempty"`
	// 请求内容格式类型：  application/json application/xml multipart/form-date text/plain  暂不支持

	ContentType *ListApiRuntimeDefinitionV2ResponseContentType `json:"content_type,omitempty"`
	// API编号

	Id *string `json:"id,omitempty"`
	// API所属分组的名称

	GroupName *string `json:"group_name,omitempty"`
	// 发布的环境名

	RunEnvName *string `json:"run_env_name,omitempty"`
	// 发布的环境id

	RunEnvId *string `json:"run_env_id,omitempty"`
	// 发布记录的编号

	PublishId *string `json:"publish_id,omitempty"`
	// 分组的二级域名

	SlDomain *string `json:"sl_domain,omitempty"`
	// 系统默认分配的子域名列表

	SlDomains *[]string `json:"sl_domains,omitempty"`
	// API的请求参数列表

	ReqParams      *[]ReqParam `json:"req_params,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListApiRuntimeDefinitionV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiRuntimeDefinitionV2Response struct{}"
	}

	return strings.Join([]string{"ListApiRuntimeDefinitionV2Response", string(data)}, " ")
}

type ListApiRuntimeDefinitionV2ResponseType struct {
	value int32
}

type ListApiRuntimeDefinitionV2ResponseTypeEnum struct {
	E_1 ListApiRuntimeDefinitionV2ResponseType
	E_2 ListApiRuntimeDefinitionV2ResponseType
}

func GetListApiRuntimeDefinitionV2ResponseTypeEnum() ListApiRuntimeDefinitionV2ResponseTypeEnum {
	return ListApiRuntimeDefinitionV2ResponseTypeEnum{
		E_1: ListApiRuntimeDefinitionV2ResponseType{
			value: 1,
		}, E_2: ListApiRuntimeDefinitionV2ResponseType{
			value: 2,
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseType) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseReqProtocol struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseReqProtocolEnum struct {
	HTTP  ListApiRuntimeDefinitionV2ResponseReqProtocol
	HTTPS ListApiRuntimeDefinitionV2ResponseReqProtocol
	BOTH  ListApiRuntimeDefinitionV2ResponseReqProtocol
}

func GetListApiRuntimeDefinitionV2ResponseReqProtocolEnum() ListApiRuntimeDefinitionV2ResponseReqProtocolEnum {
	return ListApiRuntimeDefinitionV2ResponseReqProtocolEnum{
		HTTP: ListApiRuntimeDefinitionV2ResponseReqProtocol{
			value: "HTTP",
		},
		HTTPS: ListApiRuntimeDefinitionV2ResponseReqProtocol{
			value: "HTTPS",
		},
		BOTH: ListApiRuntimeDefinitionV2ResponseReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseReqProtocol) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseReqMethod struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseReqMethodEnum struct {
	GET     ListApiRuntimeDefinitionV2ResponseReqMethod
	POST    ListApiRuntimeDefinitionV2ResponseReqMethod
	PUT     ListApiRuntimeDefinitionV2ResponseReqMethod
	DELETE  ListApiRuntimeDefinitionV2ResponseReqMethod
	HEAD    ListApiRuntimeDefinitionV2ResponseReqMethod
	PATCH   ListApiRuntimeDefinitionV2ResponseReqMethod
	OPTIONS ListApiRuntimeDefinitionV2ResponseReqMethod
	ANY     ListApiRuntimeDefinitionV2ResponseReqMethod
}

func GetListApiRuntimeDefinitionV2ResponseReqMethodEnum() ListApiRuntimeDefinitionV2ResponseReqMethodEnum {
	return ListApiRuntimeDefinitionV2ResponseReqMethodEnum{
		GET: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "GET",
		},
		POST: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "POST",
		},
		PUT: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "PUT",
		},
		DELETE: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "DELETE",
		},
		HEAD: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "HEAD",
		},
		PATCH: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "PATCH",
		},
		OPTIONS: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "OPTIONS",
		},
		ANY: ListApiRuntimeDefinitionV2ResponseReqMethod{
			value: "ANY",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseReqMethod) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseAuthType struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseAuthTypeEnum struct {
	NONE       ListApiRuntimeDefinitionV2ResponseAuthType
	APP        ListApiRuntimeDefinitionV2ResponseAuthType
	IAM        ListApiRuntimeDefinitionV2ResponseAuthType
	AUTHORIZER ListApiRuntimeDefinitionV2ResponseAuthType
}

func GetListApiRuntimeDefinitionV2ResponseAuthTypeEnum() ListApiRuntimeDefinitionV2ResponseAuthTypeEnum {
	return ListApiRuntimeDefinitionV2ResponseAuthTypeEnum{
		NONE: ListApiRuntimeDefinitionV2ResponseAuthType{
			value: "NONE",
		},
		APP: ListApiRuntimeDefinitionV2ResponseAuthType{
			value: "APP",
		},
		IAM: ListApiRuntimeDefinitionV2ResponseAuthType{
			value: "IAM",
		},
		AUTHORIZER: ListApiRuntimeDefinitionV2ResponseAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseAuthType) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseMatchMode struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseMatchModeEnum struct {
	SWA    ListApiRuntimeDefinitionV2ResponseMatchMode
	NORMAL ListApiRuntimeDefinitionV2ResponseMatchMode
}

func GetListApiRuntimeDefinitionV2ResponseMatchModeEnum() ListApiRuntimeDefinitionV2ResponseMatchModeEnum {
	return ListApiRuntimeDefinitionV2ResponseMatchModeEnum{
		SWA: ListApiRuntimeDefinitionV2ResponseMatchMode{
			value: "SWA",
		},
		NORMAL: ListApiRuntimeDefinitionV2ResponseMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseMatchMode) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseBackendType struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseBackendTypeEnum struct {
	HTTP     ListApiRuntimeDefinitionV2ResponseBackendType
	FUNCTION ListApiRuntimeDefinitionV2ResponseBackendType
	MOCK     ListApiRuntimeDefinitionV2ResponseBackendType
}

func GetListApiRuntimeDefinitionV2ResponseBackendTypeEnum() ListApiRuntimeDefinitionV2ResponseBackendTypeEnum {
	return ListApiRuntimeDefinitionV2ResponseBackendTypeEnum{
		HTTP: ListApiRuntimeDefinitionV2ResponseBackendType{
			value: "HTTP",
		},
		FUNCTION: ListApiRuntimeDefinitionV2ResponseBackendType{
			value: "FUNCTION",
		},
		MOCK: ListApiRuntimeDefinitionV2ResponseBackendType{
			value: "MOCK",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseBackendType) UnmarshalJSON(b []byte) error {
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

type ListApiRuntimeDefinitionV2ResponseContentType struct {
	value string
}

type ListApiRuntimeDefinitionV2ResponseContentTypeEnum struct {
	APPLICATION_JSON    ListApiRuntimeDefinitionV2ResponseContentType
	APPLICATION_XML     ListApiRuntimeDefinitionV2ResponseContentType
	MULTIPART_FORM_DATE ListApiRuntimeDefinitionV2ResponseContentType
	TEXT_PLAIN          ListApiRuntimeDefinitionV2ResponseContentType
}

func GetListApiRuntimeDefinitionV2ResponseContentTypeEnum() ListApiRuntimeDefinitionV2ResponseContentTypeEnum {
	return ListApiRuntimeDefinitionV2ResponseContentTypeEnum{
		APPLICATION_JSON: ListApiRuntimeDefinitionV2ResponseContentType{
			value: "application/json",
		},
		APPLICATION_XML: ListApiRuntimeDefinitionV2ResponseContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ListApiRuntimeDefinitionV2ResponseContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ListApiRuntimeDefinitionV2ResponseContentType{
			value: "text/plain",
		},
	}
}

func (c ListApiRuntimeDefinitionV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiRuntimeDefinitionV2ResponseContentType) UnmarshalJSON(b []byte) error {
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
