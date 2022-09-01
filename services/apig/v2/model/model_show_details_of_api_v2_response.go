package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowDetailsOfApiV2Response struct {

	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name" xml:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ShowDetailsOfApiV2ResponseType `json:"type" xml:"type"`

	// API的版本
	Version *string `json:"version,omitempty" xml:"version"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol ShowDetailsOfApiV2ResponseReqProtocol `json:"req_protocol" xml:"req_protocol"`

	// API的请求方式
	ReqMethod ShowDetailsOfApiV2ResponseReqMethod `json:"req_method" xml:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。
	ReqUri string `json:"req_uri" xml:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType ShowDetailsOfApiV2ResponseAuthType `json:"auth_type" xml:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty" xml:"auth_opt"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty" xml:"cors"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ShowDetailsOfApiV2ResponseMatchMode `json:"match_mode,omitempty" xml:"match_mode"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端
	BackendType ShowDetailsOfApiV2ResponseBackendType `json:"backend_type" xml:"backend_type"`

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
	ContentType *ShowDetailsOfApiV2ResponseContentType `json:"content_type,omitempty" xml:"content_type"`

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

func (o ShowDetailsOfApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiV2Response", string(data)}, " ")
}

type ShowDetailsOfApiV2ResponseType struct {
	value int32
}

type ShowDetailsOfApiV2ResponseTypeEnum struct {
	E_1 ShowDetailsOfApiV2ResponseType
	E_2 ShowDetailsOfApiV2ResponseType
}

func GetShowDetailsOfApiV2ResponseTypeEnum() ShowDetailsOfApiV2ResponseTypeEnum {
	return ShowDetailsOfApiV2ResponseTypeEnum{
		E_1: ShowDetailsOfApiV2ResponseType{
			value: 1,
		}, E_2: ShowDetailsOfApiV2ResponseType{
			value: 2,
		},
	}
}

func (c ShowDetailsOfApiV2ResponseType) Value() int32 {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseReqProtocol struct {
	value string
}

type ShowDetailsOfApiV2ResponseReqProtocolEnum struct {
	HTTP  ShowDetailsOfApiV2ResponseReqProtocol
	HTTPS ShowDetailsOfApiV2ResponseReqProtocol
	BOTH  ShowDetailsOfApiV2ResponseReqProtocol
}

func GetShowDetailsOfApiV2ResponseReqProtocolEnum() ShowDetailsOfApiV2ResponseReqProtocolEnum {
	return ShowDetailsOfApiV2ResponseReqProtocolEnum{
		HTTP: ShowDetailsOfApiV2ResponseReqProtocol{
			value: "HTTP",
		},
		HTTPS: ShowDetailsOfApiV2ResponseReqProtocol{
			value: "HTTPS",
		},
		BOTH: ShowDetailsOfApiV2ResponseReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseReqProtocol) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseReqProtocol) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseReqMethod struct {
	value string
}

type ShowDetailsOfApiV2ResponseReqMethodEnum struct {
	GET     ShowDetailsOfApiV2ResponseReqMethod
	POST    ShowDetailsOfApiV2ResponseReqMethod
	PUT     ShowDetailsOfApiV2ResponseReqMethod
	DELETE  ShowDetailsOfApiV2ResponseReqMethod
	HEAD    ShowDetailsOfApiV2ResponseReqMethod
	PATCH   ShowDetailsOfApiV2ResponseReqMethod
	OPTIONS ShowDetailsOfApiV2ResponseReqMethod
	ANY     ShowDetailsOfApiV2ResponseReqMethod
}

func GetShowDetailsOfApiV2ResponseReqMethodEnum() ShowDetailsOfApiV2ResponseReqMethodEnum {
	return ShowDetailsOfApiV2ResponseReqMethodEnum{
		GET: ShowDetailsOfApiV2ResponseReqMethod{
			value: "GET",
		},
		POST: ShowDetailsOfApiV2ResponseReqMethod{
			value: "POST",
		},
		PUT: ShowDetailsOfApiV2ResponseReqMethod{
			value: "PUT",
		},
		DELETE: ShowDetailsOfApiV2ResponseReqMethod{
			value: "DELETE",
		},
		HEAD: ShowDetailsOfApiV2ResponseReqMethod{
			value: "HEAD",
		},
		PATCH: ShowDetailsOfApiV2ResponseReqMethod{
			value: "PATCH",
		},
		OPTIONS: ShowDetailsOfApiV2ResponseReqMethod{
			value: "OPTIONS",
		},
		ANY: ShowDetailsOfApiV2ResponseReqMethod{
			value: "ANY",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseReqMethod) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseReqMethod) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseAuthType struct {
	value string
}

type ShowDetailsOfApiV2ResponseAuthTypeEnum struct {
	NONE       ShowDetailsOfApiV2ResponseAuthType
	APP        ShowDetailsOfApiV2ResponseAuthType
	IAM        ShowDetailsOfApiV2ResponseAuthType
	AUTHORIZER ShowDetailsOfApiV2ResponseAuthType
}

func GetShowDetailsOfApiV2ResponseAuthTypeEnum() ShowDetailsOfApiV2ResponseAuthTypeEnum {
	return ShowDetailsOfApiV2ResponseAuthTypeEnum{
		NONE: ShowDetailsOfApiV2ResponseAuthType{
			value: "NONE",
		},
		APP: ShowDetailsOfApiV2ResponseAuthType{
			value: "APP",
		},
		IAM: ShowDetailsOfApiV2ResponseAuthType{
			value: "IAM",
		},
		AUTHORIZER: ShowDetailsOfApiV2ResponseAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseAuthType) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseAuthType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseMatchMode struct {
	value string
}

type ShowDetailsOfApiV2ResponseMatchModeEnum struct {
	SWA    ShowDetailsOfApiV2ResponseMatchMode
	NORMAL ShowDetailsOfApiV2ResponseMatchMode
}

func GetShowDetailsOfApiV2ResponseMatchModeEnum() ShowDetailsOfApiV2ResponseMatchModeEnum {
	return ShowDetailsOfApiV2ResponseMatchModeEnum{
		SWA: ShowDetailsOfApiV2ResponseMatchMode{
			value: "SWA",
		},
		NORMAL: ShowDetailsOfApiV2ResponseMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseMatchMode) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseMatchMode) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseBackendType struct {
	value string
}

type ShowDetailsOfApiV2ResponseBackendTypeEnum struct {
	HTTP     ShowDetailsOfApiV2ResponseBackendType
	FUNCTION ShowDetailsOfApiV2ResponseBackendType
	MOCK     ShowDetailsOfApiV2ResponseBackendType
}

func GetShowDetailsOfApiV2ResponseBackendTypeEnum() ShowDetailsOfApiV2ResponseBackendTypeEnum {
	return ShowDetailsOfApiV2ResponseBackendTypeEnum{
		HTTP: ShowDetailsOfApiV2ResponseBackendType{
			value: "HTTP",
		},
		FUNCTION: ShowDetailsOfApiV2ResponseBackendType{
			value: "FUNCTION",
		},
		MOCK: ShowDetailsOfApiV2ResponseBackendType{
			value: "MOCK",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseBackendType) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseBackendType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfApiV2ResponseContentType struct {
	value string
}

type ShowDetailsOfApiV2ResponseContentTypeEnum struct {
	APPLICATION_JSON    ShowDetailsOfApiV2ResponseContentType
	APPLICATION_XML     ShowDetailsOfApiV2ResponseContentType
	MULTIPART_FORM_DATE ShowDetailsOfApiV2ResponseContentType
	TEXT_PLAIN          ShowDetailsOfApiV2ResponseContentType
}

func GetShowDetailsOfApiV2ResponseContentTypeEnum() ShowDetailsOfApiV2ResponseContentTypeEnum {
	return ShowDetailsOfApiV2ResponseContentTypeEnum{
		APPLICATION_JSON: ShowDetailsOfApiV2ResponseContentType{
			value: "application/json",
		},
		APPLICATION_XML: ShowDetailsOfApiV2ResponseContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ShowDetailsOfApiV2ResponseContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ShowDetailsOfApiV2ResponseContentType{
			value: "text/plain",
		},
	}
}

func (c ShowDetailsOfApiV2ResponseContentType) Value() string {
	return c.value
}

func (c ShowDetailsOfApiV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiV2ResponseContentType) UnmarshalJSON(b []byte) error {
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
