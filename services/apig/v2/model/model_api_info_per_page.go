package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiInfoPerPage struct {

	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ApiInfoPerPageType `json:"type"`

	// API的版本
	Version *string `json:"version,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS - GRPCS
	ReqProtocol ApiInfoPerPageReqProtocol `json:"req_protocol"`

	// API的请求方式，当API的请求协议为GRPC类型协议时请求方式固定为POST。
	ReqMethod ApiInfoPerPageReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证，当auth_type取值为AUTHORIZER时，authorizer_id字段必须传入
	AuthType ApiInfoPerPageAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiInfoPerPageMatchMode `json:"match_mode,omitempty"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流，当backend_type取值为FUNCTION时，func_info字段必须传入 - MOCK：模拟的后端，当backend_type取值为MOCK时，mock_info字段必须传入 - GRPC：grpc后端
	BackendType ApiInfoPerPageBackendType `json:"backend_type"`

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
	ContentType *ApiInfoPerPageContentType `json:"content_type,omitempty"`

	// 是否对与FunctionGraph交互场景的body进行Base64编码。仅当content_type为application/json时，可以不对body进行Base64编码。 应用场景： - 自定义认证 - 绑定断路器插件，且断路器后端降级策略为函数后端 - API后端类型为函数工作流
	IsSendFgBodyBase64 *bool `json:"is_send_fg_body_base64,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API状态   - 1： 有效
	Status *int32 `json:"status,omitempty"`

	// 是否需要编排
	ArrangeNecessary *int32 `json:"arrange_necessary,omitempty"`

	// API注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// API修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// API所属分组的版本  默认V1，其他版本暂不支持
	GroupVersion *string `json:"group_version,omitempty"`

	// 发布的环境编号  存在多个发布记录时，环境编号之间用|隔开
	RunEnvId *string `json:"run_env_id,omitempty"`

	// 发布的环境名称  存在多个发布记录时，环境名称之间用|隔开
	RunEnvName *string `json:"run_env_name,omitempty"`

	// 发布记录编号  存在多个发布记录时，发布记录编号之间用|隔开
	PublishId *string `json:"publish_id,omitempty"`

	// 发布时间  存在多个发布记录时，发布时间之间用|隔开
	PublishTime *string `json:"publish_time,omitempty"`

	// API归属的集成应用名称  暂不支持
	RomaAppName *string `json:"roma_app_name,omitempty"`

	// 当API的后端为自定义后端时，对应的自定义后端API编号  暂不支持
	LdApiId *string `json:"ld_api_id,omitempty"`

	BackendApi *BackendApi `json:"backend_api,omitempty"`

	ApiGroupInfo *ApiGroupCommonInfo `json:"api_group_info,omitempty"`

	// API的请求参数列表
	ReqParams *[]ReqParam `json:"req_params,omitempty"`
}

func (o ApiInfoPerPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiInfoPerPage struct{}"
	}

	return strings.Join([]string{"ApiInfoPerPage", string(data)}, " ")
}

type ApiInfoPerPageType struct {
	value int32
}

type ApiInfoPerPageTypeEnum struct {
	E_1 ApiInfoPerPageType
	E_2 ApiInfoPerPageType
}

func GetApiInfoPerPageTypeEnum() ApiInfoPerPageTypeEnum {
	return ApiInfoPerPageTypeEnum{
		E_1: ApiInfoPerPageType{
			value: 1,
		}, E_2: ApiInfoPerPageType{
			value: 2,
		},
	}
}

func (c ApiInfoPerPageType) Value() int32 {
	return c.value
}

func (c ApiInfoPerPageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageReqProtocol struct {
	value string
}

type ApiInfoPerPageReqProtocolEnum struct {
	HTTP  ApiInfoPerPageReqProtocol
	HTTPS ApiInfoPerPageReqProtocol
	BOTH  ApiInfoPerPageReqProtocol
	GRPCS ApiInfoPerPageReqProtocol
}

func GetApiInfoPerPageReqProtocolEnum() ApiInfoPerPageReqProtocolEnum {
	return ApiInfoPerPageReqProtocolEnum{
		HTTP: ApiInfoPerPageReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiInfoPerPageReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiInfoPerPageReqProtocol{
			value: "BOTH",
		},
		GRPCS: ApiInfoPerPageReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c ApiInfoPerPageReqProtocol) Value() string {
	return c.value
}

func (c ApiInfoPerPageReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageReqMethod struct {
	value string
}

type ApiInfoPerPageReqMethodEnum struct {
	GET     ApiInfoPerPageReqMethod
	POST    ApiInfoPerPageReqMethod
	PUT     ApiInfoPerPageReqMethod
	DELETE  ApiInfoPerPageReqMethod
	HEAD    ApiInfoPerPageReqMethod
	PATCH   ApiInfoPerPageReqMethod
	OPTIONS ApiInfoPerPageReqMethod
	ANY     ApiInfoPerPageReqMethod
}

func GetApiInfoPerPageReqMethodEnum() ApiInfoPerPageReqMethodEnum {
	return ApiInfoPerPageReqMethodEnum{
		GET: ApiInfoPerPageReqMethod{
			value: "GET",
		},
		POST: ApiInfoPerPageReqMethod{
			value: "POST",
		},
		PUT: ApiInfoPerPageReqMethod{
			value: "PUT",
		},
		DELETE: ApiInfoPerPageReqMethod{
			value: "DELETE",
		},
		HEAD: ApiInfoPerPageReqMethod{
			value: "HEAD",
		},
		PATCH: ApiInfoPerPageReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiInfoPerPageReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiInfoPerPageReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiInfoPerPageReqMethod) Value() string {
	return c.value
}

func (c ApiInfoPerPageReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageAuthType struct {
	value string
}

type ApiInfoPerPageAuthTypeEnum struct {
	NONE       ApiInfoPerPageAuthType
	APP        ApiInfoPerPageAuthType
	IAM        ApiInfoPerPageAuthType
	AUTHORIZER ApiInfoPerPageAuthType
}

func GetApiInfoPerPageAuthTypeEnum() ApiInfoPerPageAuthTypeEnum {
	return ApiInfoPerPageAuthTypeEnum{
		NONE: ApiInfoPerPageAuthType{
			value: "NONE",
		},
		APP: ApiInfoPerPageAuthType{
			value: "APP",
		},
		IAM: ApiInfoPerPageAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiInfoPerPageAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiInfoPerPageAuthType) Value() string {
	return c.value
}

func (c ApiInfoPerPageAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageAuthType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageMatchMode struct {
	value string
}

type ApiInfoPerPageMatchModeEnum struct {
	SWA    ApiInfoPerPageMatchMode
	NORMAL ApiInfoPerPageMatchMode
}

func GetApiInfoPerPageMatchModeEnum() ApiInfoPerPageMatchModeEnum {
	return ApiInfoPerPageMatchModeEnum{
		SWA: ApiInfoPerPageMatchMode{
			value: "SWA",
		},
		NORMAL: ApiInfoPerPageMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiInfoPerPageMatchMode) Value() string {
	return c.value
}

func (c ApiInfoPerPageMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageBackendType struct {
	value string
}

type ApiInfoPerPageBackendTypeEnum struct {
	HTTP     ApiInfoPerPageBackendType
	FUNCTION ApiInfoPerPageBackendType
	MOCK     ApiInfoPerPageBackendType
	GRPC     ApiInfoPerPageBackendType
}

func GetApiInfoPerPageBackendTypeEnum() ApiInfoPerPageBackendTypeEnum {
	return ApiInfoPerPageBackendTypeEnum{
		HTTP: ApiInfoPerPageBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiInfoPerPageBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiInfoPerPageBackendType{
			value: "MOCK",
		},
		GRPC: ApiInfoPerPageBackendType{
			value: "GRPC",
		},
	}
}

func (c ApiInfoPerPageBackendType) Value() string {
	return c.value
}

func (c ApiInfoPerPageBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageBackendType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageContentType struct {
	value string
}

type ApiInfoPerPageContentTypeEnum struct {
	APPLICATION_JSON    ApiInfoPerPageContentType
	APPLICATION_XML     ApiInfoPerPageContentType
	MULTIPART_FORM_DATA ApiInfoPerPageContentType
	TEXT_PLAIN          ApiInfoPerPageContentType
}

func GetApiInfoPerPageContentTypeEnum() ApiInfoPerPageContentTypeEnum {
	return ApiInfoPerPageContentTypeEnum{
		APPLICATION_JSON: ApiInfoPerPageContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiInfoPerPageContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATA: ApiInfoPerPageContentType{
			value: "multipart/form-data",
		},
		TEXT_PLAIN: ApiInfoPerPageContentType{
			value: "text/plain",
		},
	}
}

func (c ApiInfoPerPageContentType) Value() string {
	return c.value
}

func (c ApiInfoPerPageContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageContentType) UnmarshalJSON(b []byte) error {
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
