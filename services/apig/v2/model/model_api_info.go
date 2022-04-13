package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type ApiInfo struct {
	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型 - 1：公有API - 2：私有API

	Type ApiInfoType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol ApiInfoReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod ApiInfoReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType ApiInfoAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *ApiInfoMatchMode `json:"match_mode,omitempty"`
	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType ApiInfoBackendType `json:"backend_type"`
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

	ContentType *ApiInfoContentType `json:"content_type,omitempty"`
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

	PolicyHttps *[]ApiPolicyHttpResp `json:"policy_https,omitempty"`
}

func (o ApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiInfo struct{}"
	}

	return strings.Join([]string{"ApiInfo", string(data)}, " ")
}

type ApiInfoType struct {
	value int32
}

type ApiInfoTypeEnum struct {
	E_1 ApiInfoType
	E_2 ApiInfoType
}

func GetApiInfoTypeEnum() ApiInfoTypeEnum {
	return ApiInfoTypeEnum{
		E_1: ApiInfoType{
			value: 1,
		}, E_2: ApiInfoType{
			value: 2,
		},
	}
}

func (c ApiInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoType) UnmarshalJSON(b []byte) error {
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

type ApiInfoReqProtocol struct {
	value string
}

type ApiInfoReqProtocolEnum struct {
	HTTP  ApiInfoReqProtocol
	HTTPS ApiInfoReqProtocol
	BOTH  ApiInfoReqProtocol
}

func GetApiInfoReqProtocolEnum() ApiInfoReqProtocolEnum {
	return ApiInfoReqProtocolEnum{
		HTTP: ApiInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiInfoReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiInfoReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiInfoReqMethod struct {
	value string
}

type ApiInfoReqMethodEnum struct {
	GET     ApiInfoReqMethod
	POST    ApiInfoReqMethod
	PUT     ApiInfoReqMethod
	DELETE  ApiInfoReqMethod
	HEAD    ApiInfoReqMethod
	PATCH   ApiInfoReqMethod
	OPTIONS ApiInfoReqMethod
	ANY     ApiInfoReqMethod
}

func GetApiInfoReqMethodEnum() ApiInfoReqMethodEnum {
	return ApiInfoReqMethodEnum{
		GET: ApiInfoReqMethod{
			value: "GET",
		},
		POST: ApiInfoReqMethod{
			value: "POST",
		},
		PUT: ApiInfoReqMethod{
			value: "PUT",
		},
		DELETE: ApiInfoReqMethod{
			value: "DELETE",
		},
		HEAD: ApiInfoReqMethod{
			value: "HEAD",
		},
		PATCH: ApiInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiInfoAuthType struct {
	value string
}

type ApiInfoAuthTypeEnum struct {
	NONE       ApiInfoAuthType
	APP        ApiInfoAuthType
	IAM        ApiInfoAuthType
	AUTHORIZER ApiInfoAuthType
}

func GetApiInfoAuthTypeEnum() ApiInfoAuthTypeEnum {
	return ApiInfoAuthTypeEnum{
		NONE: ApiInfoAuthType{
			value: "NONE",
		},
		APP: ApiInfoAuthType{
			value: "APP",
		},
		IAM: ApiInfoAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiInfoAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoAuthType) UnmarshalJSON(b []byte) error {
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

type ApiInfoMatchMode struct {
	value string
}

type ApiInfoMatchModeEnum struct {
	SWA    ApiInfoMatchMode
	NORMAL ApiInfoMatchMode
}

func GetApiInfoMatchModeEnum() ApiInfoMatchModeEnum {
	return ApiInfoMatchModeEnum{
		SWA: ApiInfoMatchMode{
			value: "SWA",
		},
		NORMAL: ApiInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiInfoBackendType struct {
	value string
}

type ApiInfoBackendTypeEnum struct {
	HTTP     ApiInfoBackendType
	FUNCTION ApiInfoBackendType
	MOCK     ApiInfoBackendType
}

func GetApiInfoBackendTypeEnum() ApiInfoBackendTypeEnum {
	return ApiInfoBackendTypeEnum{
		HTTP: ApiInfoBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiInfoBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiInfoBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiInfoBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoBackendType) UnmarshalJSON(b []byte) error {
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

type ApiInfoContentType struct {
	value string
}

type ApiInfoContentTypeEnum struct {
	APPLICATION_JSON    ApiInfoContentType
	APPLICATION_XML     ApiInfoContentType
	MULTIPART_FORM_DATE ApiInfoContentType
	TEXT_PLAIN          ApiInfoContentType
}

func GetApiInfoContentTypeEnum() ApiInfoContentTypeEnum {
	return ApiInfoContentTypeEnum{
		APPLICATION_JSON: ApiInfoContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiInfoContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiInfoContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiInfoContentType{
			value: "text/plain",
		},
	}
}

func (c ApiInfoContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoContentType) UnmarshalJSON(b []byte) error {
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
