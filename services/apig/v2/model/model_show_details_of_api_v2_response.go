package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDetailsOfApiV2Response Response Object
type ShowDetailsOfApiV2Response struct {

	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ShowDetailsOfApiV2ResponseType `json:"type"`

	// API的版本
	Version *string `json:"version,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS - GRPCS
	ReqProtocol ShowDetailsOfApiV2ResponseReqProtocol `json:"req_protocol"`

	// API的请求方式，当API的请求协议为GRPC类型协议时请求方式固定为POST。
	ReqMethod ShowDetailsOfApiV2ResponseReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证，当auth_type取值为AUTHORIZER时，authorizer_id字段必须传入
	AuthType ShowDetailsOfApiV2ResponseAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ShowDetailsOfApiV2ResponseMatchMode `json:"match_mode,omitempty"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流，当backend_type取值为FUNCTION时，func_info字段必须传入 - MOCK：模拟的后端，当backend_type取值为MOCK时，mock_info字段必须传入 - GRPC：grpc后端
	BackendType ShowDetailsOfApiV2ResponseBackendType `json:"backend_type"`

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
	ContentType *ShowDetailsOfApiV2ResponseContentType `json:"content_type,omitempty"`

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

	FuncInfo *ApiFunc `json:"func_info,omitempty"`

	MockInfo *ApiMock `json:"mock_info,omitempty"`

	// API的请求参数列表
	ReqParams *[]ReqParam `json:"req_params,omitempty"`

	// API的后端参数列表
	BackendParams *[]BackendParam `json:"backend_params,omitempty"`

	// 函数工作流策略后端列表
	PolicyFunctions *[]ApiPolicyFunctionResp `json:"policy_functions,omitempty"`

	// mock策略后端列表
	PolicyMocks *[]ApiPolicyMockResp `json:"policy_mocks,omitempty"`

	// web策略后端列表
	PolicyHttps    *[]ApiPolicyHttpResp `json:"policy_https,omitempty"`
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

type ShowDetailsOfApiV2ResponseReqProtocol struct {
	value string
}

type ShowDetailsOfApiV2ResponseReqProtocolEnum struct {
	HTTP  ShowDetailsOfApiV2ResponseReqProtocol
	HTTPS ShowDetailsOfApiV2ResponseReqProtocol
	BOTH  ShowDetailsOfApiV2ResponseReqProtocol
	GRPCS ShowDetailsOfApiV2ResponseReqProtocol
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
		GRPCS: ShowDetailsOfApiV2ResponseReqProtocol{
			value: "GRPCS",
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

type ShowDetailsOfApiV2ResponseBackendType struct {
	value string
}

type ShowDetailsOfApiV2ResponseBackendTypeEnum struct {
	HTTP     ShowDetailsOfApiV2ResponseBackendType
	FUNCTION ShowDetailsOfApiV2ResponseBackendType
	MOCK     ShowDetailsOfApiV2ResponseBackendType
	GRPC     ShowDetailsOfApiV2ResponseBackendType
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
		GRPC: ShowDetailsOfApiV2ResponseBackendType{
			value: "GRPC",
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

type ShowDetailsOfApiV2ResponseContentType struct {
	value string
}

type ShowDetailsOfApiV2ResponseContentTypeEnum struct {
	APPLICATION_JSON    ShowDetailsOfApiV2ResponseContentType
	APPLICATION_XML     ShowDetailsOfApiV2ResponseContentType
	MULTIPART_FORM_DATA ShowDetailsOfApiV2ResponseContentType
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
		MULTIPART_FORM_DATA: ShowDetailsOfApiV2ResponseContentType{
			value: "multipart/form-data",
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
