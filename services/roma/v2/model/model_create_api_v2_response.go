package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// Response Object
type CreateApiV2Response struct {
	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型[，该参数暂未使用](tag:hcs;fcs;) - 1：公有API - 2：私有API

	Type CreateApiV2ResponseType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol CreateApiV2ResponseReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod CreateApiV2ResponseReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式[，site暂不支持IAM认证。](tag:Site) - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType CreateApiV2ResponseAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *CreateApiV2ResponseMatchMode `json:"match_mode,omitempty"`
	// 后端类型[，site暂不支持函数工作流。](tag:Site) - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType CreateApiV2ResponseBackendType `json:"backend_type"`
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

	ContentType *CreateApiV2ResponseContentType `json:"content_type,omitempty"`
	// API编号

	Id *string `json:"id,omitempty"`
	// API状态   - 1： 有效

	Status *CreateApiV2ResponseStatus `json:"status,omitempty"`
	// 是否需要编排

	ArrangeNecessary *int32 `json:"arrange_necessary,omitempty"`
	// API注册时间

	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`
	// API修改时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// API所属分组的名称

	GroupName *string `json:"group_name,omitempty"`
	// API所属分组的版本

	GroupVersion *string `json:"group_version,omitempty"`
	// 发布的环境名称  存在多个发布记录时，环境名称之间用|隔开

	RunEnvName *string `json:"run_env_name,omitempty"`
	// 发布的环境编号  存在多个发布记录时，环境编号之间用|隔开

	RunEnvId *string `json:"run_env_id,omitempty"`
	// 发布记录编号  存在多个发布记录时，发布记录编号之间用|隔开

	PublishId *string `json:"publish_id,omitempty"`
	// 发布时间  存在多个发布记录时，发布时间之间用|隔开

	PublishTime *string `json:"publish_time,omitempty"`
	// API归属的集成应用名称

	RomaAppName *string `json:"roma_app_name,omitempty"`
	// 当API的后端为自定义后端时，对应的自定义后端API编号

	LdApiId *string `json:"ld_api_id,omitempty"`

	BackendApi *BackendApi `json:"backend_api,omitempty"`

	ApiGroupInfo *ApiGroupCommonInfo `json:"api_group_info,omitempty"`

	FuncInfo *ApiFunc `json:"func_info,omitempty"`

	MockInfo *ApiMock `json:"mock_info,omitempty"`
	// API的请求参数列表

	ReqParams *[]ReqParam `json:"req_params,omitempty"`
	// API的后端参数列表

	BackendParams *[]BackendParam `json:"backend_params,omitempty"`
	// [函数工作流策略后端列表](tag:hws;hws_hk;hcs;fcs;g42;)[暂不支持](tag:Site)

	PolicyFunctions *[]ApiPolicyFunctionResp `json:"policy_functions,omitempty"`
	// mock策略后端列表

	PolicyMocks *[]ApiPolicyMockResp `json:"policy_mocks,omitempty"`
	// web策略后端列表

	PolicyHttps    *[]ApiPolicyHttpResp `json:"policy_https,omitempty"`
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

type CreateApiV2ResponseStatus struct {
	value int32
}

type CreateApiV2ResponseStatusEnum struct {
	E_1 CreateApiV2ResponseStatus
}

func GetCreateApiV2ResponseStatusEnum() CreateApiV2ResponseStatusEnum {
	return CreateApiV2ResponseStatusEnum{
		E_1: CreateApiV2ResponseStatus{
			value: 1,
		},
	}
}

func (c CreateApiV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
