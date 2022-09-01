package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateApiV2Response struct {

	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name" xml:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type CreateApiV2ResponseType `json:"type" xml:"type"`

	// API的版本
	Version *string `json:"version,omitempty" xml:"version"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol CreateApiV2ResponseReqProtocol `json:"req_protocol" xml:"req_protocol"`

	// API的请求方式
	ReqMethod CreateApiV2ResponseReqMethod `json:"req_method" xml:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。
	ReqUri string `json:"req_uri" xml:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType CreateApiV2ResponseAuthType `json:"auth_type" xml:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty" xml:"auth_opt"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty" xml:"cors"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *CreateApiV2ResponseMatchMode `json:"match_mode,omitempty" xml:"match_mode"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端
	BackendType CreateApiV2ResponseBackendType `json:"backend_type" xml:"backend_type"`

	// API描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// API所属的分组编号
	GroupId string `json:"group_id" xml:"group_id"`

	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	BodyRemark *string `json:"body_remark,omitempty" xml:"body_remark"`

	// 正常响应示例，描述API的正常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultNormalSample *string `json:"result_normal_sample,omitempty" xml:"result_normal_sample"`

	// 失败返回示例，描述API的异常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultFailureSample *string `json:"result_failure_sample,omitempty" xml:"result_failure_sample"`

	// 前端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty" xml:"authorizer_id"`

	// 标签。  支持英文，数字，下划线，且只能以英文开头。支持输入多个标签，不同标签以英文逗号分割。
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 分组自定义响应ID
	ResponseId *string `json:"response_id,omitempty" xml:"response_id"`

	// 集成应用ID  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty" xml:"roma_app_id"`

	// API绑定的自定义域名  暂不支持
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 标签  待废弃，优先使用tags字段
	Tag *string `json:"tag,omitempty" xml:"tag"`

	// 请求内容格式类型：  application/json application/xml multipart/form-date text/plain  暂不支持
	ContentType *CreateApiV2ResponseContentType `json:"content_type,omitempty" xml:"content_type"`

	// API编号
	Id *string `json:"id,omitempty" xml:"id"`

	// API状态   - 1： 有效
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 是否需要编排
	ArrangeNecessary *int32 `json:"arrange_necessary,omitempty" xml:"arrange_necessary"`

	// API注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty" xml:"register_time"`

	// API修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty" xml:"update_time"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// API所属分组的版本  默认V1，其他版本暂不支持
	GroupVersion *string `json:"group_version,omitempty" xml:"group_version"`

	// 发布的环境编号  存在多个发布记录时，环境编号之间用|隔开
	RunEnvId *string `json:"run_env_id,omitempty" xml:"run_env_id"`

	// 发布的环境名称  存在多个发布记录时，环境名称之间用|隔开
	RunEnvName *string `json:"run_env_name,omitempty" xml:"run_env_name"`

	// 发布记录编号  存在多个发布记录时，发布记录编号之间用|隔开
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// 发布时间  存在多个发布记录时，发布时间之间用|隔开
	PublishTime *string `json:"publish_time,omitempty" xml:"publish_time"`

	// API归属的集成应用名称  暂不支持
	RomaAppName *string `json:"roma_app_name,omitempty" xml:"roma_app_name"`

	// 当API的后端为自定义后端时，对应的自定义后端API编号  暂不支持
	LdApiId *string `json:"ld_api_id,omitempty" xml:"ld_api_id"`

	BackendApi *BackendApi `json:"backend_api,omitempty" xml:"backend_api"`

	ApiGroupInfo *ApiGroupCommonInfo `json:"api_group_info,omitempty" xml:"api_group_info"`

	FuncInfo *ApiFunc `json:"func_info,omitempty" xml:"func_info"`

	MockInfo *ApiMock `json:"mock_info,omitempty" xml:"mock_info"`

	// API的请求参数列表
	ReqParams *[]ReqParam `json:"req_params,omitempty" xml:"req_params"`

	// API的后端参数列表
	BackendParams *[]BackendParam `json:"backend_params,omitempty" xml:"backend_params"`

	// 函数工作流策略后端列表
	PolicyFunctions *[]ApiPolicyFunctionResp `json:"policy_functions,omitempty" xml:"policy_functions"`

	// mock策略后端列表
	PolicyMocks *[]ApiPolicyMockResp `json:"policy_mocks,omitempty" xml:"policy_mocks"`

	// web策略后端列表
	PolicyHttps    *[]ApiPolicyHttpResp `json:"policy_https,omitempty" xml:"policy_https"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiV2Response struct{}"
	}

	return strings.Join([]string{"CreateApiV2Response", string(data)}, " ")
}

type CreateApiV2ResponseType struct {
	value int32
}

type CreateApiV2ResponseTypeEnum struct {
	E_1 CreateApiV2ResponseType
	E_2 CreateApiV2ResponseType
}

func GetCreateApiV2ResponseTypeEnum() CreateApiV2ResponseTypeEnum {
	return CreateApiV2ResponseTypeEnum{
		E_1: CreateApiV2ResponseType{
			value: 1,
		}, E_2: CreateApiV2ResponseType{
			value: 2,
		},
	}
}

func (c CreateApiV2ResponseType) Value() int32 {
	return c.value
}

func (c CreateApiV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseType) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseReqProtocol struct {
	value string
}

type CreateApiV2ResponseReqProtocolEnum struct {
	HTTP  CreateApiV2ResponseReqProtocol
	HTTPS CreateApiV2ResponseReqProtocol
	BOTH  CreateApiV2ResponseReqProtocol
}

func GetCreateApiV2ResponseReqProtocolEnum() CreateApiV2ResponseReqProtocolEnum {
	return CreateApiV2ResponseReqProtocolEnum{
		HTTP: CreateApiV2ResponseReqProtocol{
			value: "HTTP",
		},
		HTTPS: CreateApiV2ResponseReqProtocol{
			value: "HTTPS",
		},
		BOTH: CreateApiV2ResponseReqProtocol{
			value: "BOTH",
		},
	}
}

func (c CreateApiV2ResponseReqProtocol) Value() string {
	return c.value
}

func (c CreateApiV2ResponseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseReqProtocol) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseReqMethod struct {
	value string
}

type CreateApiV2ResponseReqMethodEnum struct {
	GET     CreateApiV2ResponseReqMethod
	POST    CreateApiV2ResponseReqMethod
	PUT     CreateApiV2ResponseReqMethod
	DELETE  CreateApiV2ResponseReqMethod
	HEAD    CreateApiV2ResponseReqMethod
	PATCH   CreateApiV2ResponseReqMethod
	OPTIONS CreateApiV2ResponseReqMethod
	ANY     CreateApiV2ResponseReqMethod
}

func GetCreateApiV2ResponseReqMethodEnum() CreateApiV2ResponseReqMethodEnum {
	return CreateApiV2ResponseReqMethodEnum{
		GET: CreateApiV2ResponseReqMethod{
			value: "GET",
		},
		POST: CreateApiV2ResponseReqMethod{
			value: "POST",
		},
		PUT: CreateApiV2ResponseReqMethod{
			value: "PUT",
		},
		DELETE: CreateApiV2ResponseReqMethod{
			value: "DELETE",
		},
		HEAD: CreateApiV2ResponseReqMethod{
			value: "HEAD",
		},
		PATCH: CreateApiV2ResponseReqMethod{
			value: "PATCH",
		},
		OPTIONS: CreateApiV2ResponseReqMethod{
			value: "OPTIONS",
		},
		ANY: CreateApiV2ResponseReqMethod{
			value: "ANY",
		},
	}
}

func (c CreateApiV2ResponseReqMethod) Value() string {
	return c.value
}

func (c CreateApiV2ResponseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseReqMethod) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseAuthType struct {
	value string
}

type CreateApiV2ResponseAuthTypeEnum struct {
	NONE       CreateApiV2ResponseAuthType
	APP        CreateApiV2ResponseAuthType
	IAM        CreateApiV2ResponseAuthType
	AUTHORIZER CreateApiV2ResponseAuthType
}

func GetCreateApiV2ResponseAuthTypeEnum() CreateApiV2ResponseAuthTypeEnum {
	return CreateApiV2ResponseAuthTypeEnum{
		NONE: CreateApiV2ResponseAuthType{
			value: "NONE",
		},
		APP: CreateApiV2ResponseAuthType{
			value: "APP",
		},
		IAM: CreateApiV2ResponseAuthType{
			value: "IAM",
		},
		AUTHORIZER: CreateApiV2ResponseAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c CreateApiV2ResponseAuthType) Value() string {
	return c.value
}

func (c CreateApiV2ResponseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseAuthType) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseMatchMode struct {
	value string
}

type CreateApiV2ResponseMatchModeEnum struct {
	SWA    CreateApiV2ResponseMatchMode
	NORMAL CreateApiV2ResponseMatchMode
}

func GetCreateApiV2ResponseMatchModeEnum() CreateApiV2ResponseMatchModeEnum {
	return CreateApiV2ResponseMatchModeEnum{
		SWA: CreateApiV2ResponseMatchMode{
			value: "SWA",
		},
		NORMAL: CreateApiV2ResponseMatchMode{
			value: "NORMAL",
		},
	}
}

func (c CreateApiV2ResponseMatchMode) Value() string {
	return c.value
}

func (c CreateApiV2ResponseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseMatchMode) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseBackendType struct {
	value string
}

type CreateApiV2ResponseBackendTypeEnum struct {
	HTTP     CreateApiV2ResponseBackendType
	FUNCTION CreateApiV2ResponseBackendType
	MOCK     CreateApiV2ResponseBackendType
}

func GetCreateApiV2ResponseBackendTypeEnum() CreateApiV2ResponseBackendTypeEnum {
	return CreateApiV2ResponseBackendTypeEnum{
		HTTP: CreateApiV2ResponseBackendType{
			value: "HTTP",
		},
		FUNCTION: CreateApiV2ResponseBackendType{
			value: "FUNCTION",
		},
		MOCK: CreateApiV2ResponseBackendType{
			value: "MOCK",
		},
	}
}

func (c CreateApiV2ResponseBackendType) Value() string {
	return c.value
}

func (c CreateApiV2ResponseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseBackendType) UnmarshalJSON(b []byte) error {
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

type CreateApiV2ResponseContentType struct {
	value string
}

type CreateApiV2ResponseContentTypeEnum struct {
	APPLICATION_JSON    CreateApiV2ResponseContentType
	APPLICATION_XML     CreateApiV2ResponseContentType
	MULTIPART_FORM_DATE CreateApiV2ResponseContentType
	TEXT_PLAIN          CreateApiV2ResponseContentType
}

func GetCreateApiV2ResponseContentTypeEnum() CreateApiV2ResponseContentTypeEnum {
	return CreateApiV2ResponseContentTypeEnum{
		APPLICATION_JSON: CreateApiV2ResponseContentType{
			value: "application/json",
		},
		APPLICATION_XML: CreateApiV2ResponseContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: CreateApiV2ResponseContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: CreateApiV2ResponseContentType{
			value: "text/plain",
		},
	}
}

func (c CreateApiV2ResponseContentType) Value() string {
	return c.value
}

func (c CreateApiV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseContentType) UnmarshalJSON(b []byte) error {
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
