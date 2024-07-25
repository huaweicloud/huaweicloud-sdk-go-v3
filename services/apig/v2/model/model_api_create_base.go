package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiCreateBase struct {

	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ApiCreateBaseType `json:"type"`

	// API的版本
	Version *string `json:"version,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS - GRPCS
	ReqProtocol ApiCreateBaseReqProtocol `json:"req_protocol"`

	// API的请求方式，当API的请求协议为GRPC类型协议时请求方式固定为POST。
	ReqMethod ApiCreateBaseReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证，当auth_type取值为AUTHORIZER时，authorizer_id字段必须传入
	AuthType ApiCreateBaseAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiCreateBaseMatchMode `json:"match_mode,omitempty"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流，当backend_type取值为FUNCTION时，func_info字段必须传入 - MOCK：模拟的后端，当backend_type取值为MOCK时，mock_info字段必须传入 - GRPC：grpc后端
	BackendType ApiCreateBaseBackendType `json:"backend_type"`

	// API描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// API所属的分组编号
	GroupId string `json:"group_id"`

	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	BodyRemark *string `json:"body_remark,omitempty"`

	// 正常响应示例，描述API的正常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。  当API的请求协议为GRPC类型时不支持配置。
	ResultNormalSample *string `json:"result_normal_sample,omitempty"`

	// 失败返回示例，描述API的异常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。  当API的请求协议为GRPC类型时不支持配置。
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

	// 请求内容格式类型：  application/json application/xml multipart/form-data text/plain
	ContentType *ApiCreateBaseContentType `json:"content_type,omitempty"`

	// 是否对与FunctionGraph交互场景的body进行Base64编码。仅当content_type为application/json时，可以不对body进行Base64编码。 应用场景： - 自定义认证 - 绑定断路器插件，且断路器后端降级策略为函数后端 - API后端类型为函数工作流
	IsSendFgBodyBase64 *bool `json:"is_send_fg_body_base64,omitempty"`

	MockInfo *ApiMockCreate `json:"mock_info,omitempty"`

	FuncInfo *ApiFuncCreate `json:"func_info,omitempty"`

	// API的请求参数列表，API请求协议为GRPC类型时不支持配置
	ReqParams *[]ReqParamBase `json:"req_params,omitempty"`

	// API的后端参数列表，API请求协议为GRPC类型时不支持配置
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`

	// mock策略后端列表
	PolicyMocks *[]ApiPolicyMockCreate `json:"policy_mocks,omitempty"`

	// 函数工作流策略后端列表
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

func (c ApiCreateBaseType) Value() int32 {
	return c.value
}

func (c ApiCreateBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseReqProtocol struct {
	value string
}

type ApiCreateBaseReqProtocolEnum struct {
	HTTP  ApiCreateBaseReqProtocol
	HTTPS ApiCreateBaseReqProtocol
	BOTH  ApiCreateBaseReqProtocol
	GRPCS ApiCreateBaseReqProtocol
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
		GRPCS: ApiCreateBaseReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiCreateBaseReqProtocol) Value() string {
	return c.value
}

func (c ApiCreateBaseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqProtocol) UnmarshalJSON(b []byte) error {
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

func (c ApiCreateBaseReqMethod) Value() string {
	return c.value
}

func (c ApiCreateBaseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqMethod) UnmarshalJSON(b []byte) error {
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

func (c ApiCreateBaseAuthType) Value() string {
	return c.value
}

func (c ApiCreateBaseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseAuthType) UnmarshalJSON(b []byte) error {
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

func (c ApiCreateBaseMatchMode) Value() string {
	return c.value
}

func (c ApiCreateBaseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseBackendType struct {
	value string
}

type ApiCreateBaseBackendTypeEnum struct {
	HTTP     ApiCreateBaseBackendType
	FUNCTION ApiCreateBaseBackendType
	MOCK     ApiCreateBaseBackendType
	GRPC     ApiCreateBaseBackendType
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
		GRPC: ApiCreateBaseBackendType{
			value: "GRPC",
		},
	}
}

func (c ApiCreateBaseBackendType) Value() string {
	return c.value
}

func (c ApiCreateBaseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseBackendType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseContentType struct {
	value string
}

type ApiCreateBaseContentTypeEnum struct {
	APPLICATION_JSON    ApiCreateBaseContentType
	APPLICATION_XML     ApiCreateBaseContentType
	MULTIPART_FORM_DATA ApiCreateBaseContentType
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
		MULTIPART_FORM_DATA: ApiCreateBaseContentType{
			value: "multipart/form-data",
		},
		TEXT_PLAIN: ApiCreateBaseContentType{
			value: "text/plain",
		},
	}
}

func (c ApiCreateBaseContentType) Value() string {
	return c.value
}

func (c ApiCreateBaseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseContentType) UnmarshalJSON(b []byte) error {
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
