package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiCreate struct {

	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ApiCreateType `json:"type"`

	// API的版本
	Version *string `json:"version,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol ApiCreateReqProtocol `json:"req_protocol"`

	// API的请求方式
	ReqMethod ApiCreateReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType ApiCreateAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiCreateMatchMode `json:"match_mode,omitempty"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端
	BackendType ApiCreateBackendType `json:"backend_type"`

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

	// 标签。  支持英文，数字，中文，特殊符号（-*#%.:_），且只能以中文或英文开头。  默认支持10个标签，如需扩大配额请联系技术工程师修改API_TAG_NUM_LIMIT配置。
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
	ContentType *ApiCreateContentType `json:"content_type,omitempty"`

	MockInfo *ApiMockCreate `json:"mock_info,omitempty"`

	FuncInfo *ApiFuncCreate `json:"func_info,omitempty"`

	// API的请求参数列表
	ReqParams *[]ReqParamBase `json:"req_params,omitempty"`

	// API的后端参数列表
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`

	// mock策略后端列表
	PolicyMocks *[]ApiPolicyMockCreate `json:"policy_mocks,omitempty"`

	// 函数工作流策略后端列表
	PolicyFunctions *[]ApiPolicyFunctionCreate `json:"policy_functions,omitempty"`

	BackendApi *BackendApiCreate `json:"backend_api,omitempty"`

	// web策略后端列表
	PolicyHttps *[]ApiPolicyHttpCreate `json:"policy_https,omitempty"`
}

func (o ApiCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCreate struct{}"
	}

	return strings.Join([]string{"ApiCreate", string(data)}, " ")
}

type ApiCreateType struct {
	value int32
}

type ApiCreateTypeEnum struct {
	E_1 ApiCreateType
	E_2 ApiCreateType
}

func GetApiCreateTypeEnum() ApiCreateTypeEnum {
	return ApiCreateTypeEnum{
		E_1: ApiCreateType{
			value: 1,
		}, E_2: ApiCreateType{
			value: 2,
		},
	}
}

func (c ApiCreateType) Value() int32 {
	return c.value
}

func (c ApiCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ApiCreateReqProtocol struct {
	value string
}

type ApiCreateReqProtocolEnum struct {
	HTTP  ApiCreateReqProtocol
	HTTPS ApiCreateReqProtocol
	BOTH  ApiCreateReqProtocol
}

func GetApiCreateReqProtocolEnum() ApiCreateReqProtocolEnum {
	return ApiCreateReqProtocolEnum{
		HTTP: ApiCreateReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiCreateReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiCreateReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiCreateReqProtocol) Value() string {
	return c.value
}

func (c ApiCreateReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiCreateReqMethod struct {
	value string
}

type ApiCreateReqMethodEnum struct {
	GET     ApiCreateReqMethod
	POST    ApiCreateReqMethod
	PUT     ApiCreateReqMethod
	DELETE  ApiCreateReqMethod
	HEAD    ApiCreateReqMethod
	PATCH   ApiCreateReqMethod
	OPTIONS ApiCreateReqMethod
	ANY     ApiCreateReqMethod
}

func GetApiCreateReqMethodEnum() ApiCreateReqMethodEnum {
	return ApiCreateReqMethodEnum{
		GET: ApiCreateReqMethod{
			value: "GET",
		},
		POST: ApiCreateReqMethod{
			value: "POST",
		},
		PUT: ApiCreateReqMethod{
			value: "PUT",
		},
		DELETE: ApiCreateReqMethod{
			value: "DELETE",
		},
		HEAD: ApiCreateReqMethod{
			value: "HEAD",
		},
		PATCH: ApiCreateReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiCreateReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiCreateReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiCreateReqMethod) Value() string {
	return c.value
}

func (c ApiCreateReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiCreateAuthType struct {
	value string
}

type ApiCreateAuthTypeEnum struct {
	NONE       ApiCreateAuthType
	APP        ApiCreateAuthType
	IAM        ApiCreateAuthType
	AUTHORIZER ApiCreateAuthType
}

func GetApiCreateAuthTypeEnum() ApiCreateAuthTypeEnum {
	return ApiCreateAuthTypeEnum{
		NONE: ApiCreateAuthType{
			value: "NONE",
		},
		APP: ApiCreateAuthType{
			value: "APP",
		},
		IAM: ApiCreateAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiCreateAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiCreateAuthType) Value() string {
	return c.value
}

func (c ApiCreateAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateAuthType) UnmarshalJSON(b []byte) error {
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

type ApiCreateMatchMode struct {
	value string
}

type ApiCreateMatchModeEnum struct {
	SWA    ApiCreateMatchMode
	NORMAL ApiCreateMatchMode
}

func GetApiCreateMatchModeEnum() ApiCreateMatchModeEnum {
	return ApiCreateMatchModeEnum{
		SWA: ApiCreateMatchMode{
			value: "SWA",
		},
		NORMAL: ApiCreateMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiCreateMatchMode) Value() string {
	return c.value
}

func (c ApiCreateMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiCreateBackendType struct {
	value string
}

type ApiCreateBackendTypeEnum struct {
	HTTP     ApiCreateBackendType
	FUNCTION ApiCreateBackendType
	MOCK     ApiCreateBackendType
}

func GetApiCreateBackendTypeEnum() ApiCreateBackendTypeEnum {
	return ApiCreateBackendTypeEnum{
		HTTP: ApiCreateBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiCreateBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiCreateBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiCreateBackendType) Value() string {
	return c.value
}

func (c ApiCreateBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBackendType) UnmarshalJSON(b []byte) error {
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

type ApiCreateContentType struct {
	value string
}

type ApiCreateContentTypeEnum struct {
	APPLICATION_JSON    ApiCreateContentType
	APPLICATION_XML     ApiCreateContentType
	MULTIPART_FORM_DATE ApiCreateContentType
	TEXT_PLAIN          ApiCreateContentType
}

func GetApiCreateContentTypeEnum() ApiCreateContentTypeEnum {
	return ApiCreateContentTypeEnum{
		APPLICATION_JSON: ApiCreateContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiCreateContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiCreateContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiCreateContentType{
			value: "text/plain",
		},
	}
}

func (c ApiCreateContentType) Value() string {
	return c.value
}

func (c ApiCreateContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateContentType) UnmarshalJSON(b []byte) error {
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
