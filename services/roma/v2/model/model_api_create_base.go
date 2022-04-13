package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiCreateBase struct {
	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型[，该参数暂未使用](tag:hcs;fcs;) - 1：公有API - 2：私有API

	Type ApiCreateBaseType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol ApiCreateBaseReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod ApiCreateBaseReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式[，site暂不支持IAM认证。](tag:Site) - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType ApiCreateBaseAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *ApiCreateBaseMatchMode `json:"match_mode,omitempty"`
	// 后端类型[，site暂不支持函数工作流。](tag:Site) - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType ApiCreateBaseBackendType `json:"backend_type"`
	// API描述。  不允许带有<、>字符 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// API所属的分组编号

	GroupId string `json:"group_id"`
	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。 > 中文字符必须为UTF-8或者unicode编码。

	BodyRemark *string `json:"body_remark,omitempty"`
	// 正常响应示例，描述API的正常返回信息。 > 中文字符必须为UTF-8或者unicode编码。

	ResultNormalSample *string `json:"result_normal_sample,omitempty"`
	// 失败返回示例，描述API的异常返回信息。 > 中文字符必须为UTF-8或者unicode编码。

	ResultFailureSample *string `json:"result_failure_sample,omitempty"`
	// 前端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 标签。  支持英文，数字，中文，特殊符号（-*#%.:_），且只能以中文或英文开头。支持输入多个标签，不同标签以英文逗号分割。  默认支持10个标签，如需扩大配额请联系技术工程师修改API_TAG_NUM_LIMIT配置。

	Tags *[]string `json:"tags,omitempty"`
	// 分组自定义响应ID  暂不支持

	ResponseId *string `json:"response_id,omitempty"`
	// API归属的集成应用编号  API分组为全局分组时或API绑定自定义域名时必填。

	RomaAppId *string `json:"roma_app_id,omitempty"`
	// API绑定的自定义域名，使用自定义域名时roma_app_id字段必填。

	DomainName *string `json:"domain_name,omitempty"`
	// 标签  待废弃，优先使用tags字段

	Tag *string `json:"tag,omitempty"`
	// 请求内容格式类型：  application/json application/xml multipart/form-date text/plain

	ContentType *ApiCreateBaseContentType `json:"content_type,omitempty"`

	MockInfo *ApiMockCreate `json:"mock_info,omitempty"`

	FuncInfo *ApiFuncCreate `json:"func_info,omitempty"`
	// API的请求参数列表

	ReqParams *[]ReqParamBase `json:"req_params,omitempty"`
	// API的后端参数列表

	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`
	// mock策略后端列表

	PolicyMocks *[]ApiPolicyMockCreate `json:"policy_mocks,omitempty"`
	// [函数工作流策略后端列表](tag:hws;hws_hk;hcs;fcs;g42;)[暂不支持](tag:Site)

	PolicyFunctions *[]ApiPolicyFunctionCreate `json:"policy_functions,omitempty"`
}

func (o ApiCreateBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCreateBase struct{}"
	}

	return strings.Join([]string{"ApiCreateBase", string(data)}, " ")
}

type ApiCreateBaseType struct {
	value int32
}

type ApiCreateBaseTypeEnum struct {
	E_1 ApiCreateBaseType
	E_2 ApiCreateBaseType
}

func GetApiCreateBaseTypeEnum() ApiCreateBaseTypeEnum {
	return ApiCreateBaseTypeEnum{
		E_1: ApiCreateBaseType{
			value: 1,
		}, E_2: ApiCreateBaseType{
			value: 2,
		},
	}
}

func (c ApiCreateBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseReqProtocol struct {
	value string
}

type ApiCreateBaseReqProtocolEnum struct {
	HTTP  ApiCreateBaseReqProtocol
	HTTPS ApiCreateBaseReqProtocol
	BOTH  ApiCreateBaseReqProtocol
}

func GetApiCreateBaseReqProtocolEnum() ApiCreateBaseReqProtocolEnum {
	return ApiCreateBaseReqProtocolEnum{
		HTTP: ApiCreateBaseReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiCreateBaseReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiCreateBaseReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiCreateBaseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseReqMethod struct {
	value string
}

type ApiCreateBaseReqMethodEnum struct {
	GET     ApiCreateBaseReqMethod
	POST    ApiCreateBaseReqMethod
	PUT     ApiCreateBaseReqMethod
	DELETE  ApiCreateBaseReqMethod
	HEAD    ApiCreateBaseReqMethod
	PATCH   ApiCreateBaseReqMethod
	OPTIONS ApiCreateBaseReqMethod
	ANY     ApiCreateBaseReqMethod
}

func GetApiCreateBaseReqMethodEnum() ApiCreateBaseReqMethodEnum {
	return ApiCreateBaseReqMethodEnum{
		GET: ApiCreateBaseReqMethod{
			value: "GET",
		},
		POST: ApiCreateBaseReqMethod{
			value: "POST",
		},
		PUT: ApiCreateBaseReqMethod{
			value: "PUT",
		},
		DELETE: ApiCreateBaseReqMethod{
			value: "DELETE",
		},
		HEAD: ApiCreateBaseReqMethod{
			value: "HEAD",
		},
		PATCH: ApiCreateBaseReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiCreateBaseReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiCreateBaseReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiCreateBaseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseAuthType struct {
	value string
}

type ApiCreateBaseAuthTypeEnum struct {
	NONE       ApiCreateBaseAuthType
	APP        ApiCreateBaseAuthType
	IAM        ApiCreateBaseAuthType
	AUTHORIZER ApiCreateBaseAuthType
}

func GetApiCreateBaseAuthTypeEnum() ApiCreateBaseAuthTypeEnum {
	return ApiCreateBaseAuthTypeEnum{
		NONE: ApiCreateBaseAuthType{
			value: "NONE",
		},
		APP: ApiCreateBaseAuthType{
			value: "APP",
		},
		IAM: ApiCreateBaseAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiCreateBaseAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiCreateBaseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseAuthType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseMatchMode struct {
	value string
}

type ApiCreateBaseMatchModeEnum struct {
	SWA    ApiCreateBaseMatchMode
	NORMAL ApiCreateBaseMatchMode
}

func GetApiCreateBaseMatchModeEnum() ApiCreateBaseMatchModeEnum {
	return ApiCreateBaseMatchModeEnum{
		SWA: ApiCreateBaseMatchMode{
			value: "SWA",
		},
		NORMAL: ApiCreateBaseMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiCreateBaseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseBackendType struct {
	value string
}

type ApiCreateBaseBackendTypeEnum struct {
	HTTP     ApiCreateBaseBackendType
	FUNCTION ApiCreateBaseBackendType
	MOCK     ApiCreateBaseBackendType
}

func GetApiCreateBaseBackendTypeEnum() ApiCreateBaseBackendTypeEnum {
	return ApiCreateBaseBackendTypeEnum{
		HTTP: ApiCreateBaseBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiCreateBaseBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiCreateBaseBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiCreateBaseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseBackendType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseContentType struct {
	value string
}

type ApiCreateBaseContentTypeEnum struct {
	APPLICATION_JSON    ApiCreateBaseContentType
	APPLICATION_XML     ApiCreateBaseContentType
	MULTIPART_FORM_DATE ApiCreateBaseContentType
	TEXT_PLAIN          ApiCreateBaseContentType
}

func GetApiCreateBaseContentTypeEnum() ApiCreateBaseContentTypeEnum {
	return ApiCreateBaseContentTypeEnum{
		APPLICATION_JSON: ApiCreateBaseContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiCreateBaseContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiCreateBaseContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiCreateBaseContentType{
			value: "text/plain",
		},
	}
}

func (c ApiCreateBaseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseContentType) UnmarshalJSON(b []byte) error {
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
