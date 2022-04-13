package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type ApiInfoPerPage struct {
	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型[，该参数暂未使用](tag:hcs;fcs;) - 1：公有API - 2：私有API

	Type ApiInfoPerPageType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol ApiInfoPerPageReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod ApiInfoPerPageReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式[，site暂不支持IAM认证。](tag:Site) - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType ApiInfoPerPageAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *ApiInfoPerPageMatchMode `json:"match_mode,omitempty"`
	// 后端类型[，site暂不支持函数工作流。](tag:Site) - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType ApiInfoPerPageBackendType `json:"backend_type"`
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

	ContentType *ApiInfoPerPageContentType `json:"content_type,omitempty"`
	// API编号

	Id *string `json:"id,omitempty"`
	// API状态   - 1： 有效

	Status *ApiInfoPerPageStatus `json:"status,omitempty"`
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

func (c ApiInfoPerPageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageReqProtocol struct {
	value string
}

type ApiInfoPerPageReqProtocolEnum struct {
	HTTP  ApiInfoPerPageReqProtocol
	HTTPS ApiInfoPerPageReqProtocol
	BOTH  ApiInfoPerPageReqProtocol
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
	}
}

func (c ApiInfoPerPageReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageReqProtocol) UnmarshalJSON(b []byte) error {
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

func (c ApiInfoPerPageReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageReqMethod) UnmarshalJSON(b []byte) error {
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

func (c ApiInfoPerPageAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageAuthType) UnmarshalJSON(b []byte) error {
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

func (c ApiInfoPerPageMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageBackendType struct {
	value string
}

type ApiInfoPerPageBackendTypeEnum struct {
	HTTP     ApiInfoPerPageBackendType
	FUNCTION ApiInfoPerPageBackendType
	MOCK     ApiInfoPerPageBackendType
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
	}
}

func (c ApiInfoPerPageBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageBackendType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageContentType struct {
	value string
}

type ApiInfoPerPageContentTypeEnum struct {
	APPLICATION_JSON    ApiInfoPerPageContentType
	APPLICATION_XML     ApiInfoPerPageContentType
	MULTIPART_FORM_DATE ApiInfoPerPageContentType
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
		MULTIPART_FORM_DATE: ApiInfoPerPageContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiInfoPerPageContentType{
			value: "text/plain",
		},
	}
}

func (c ApiInfoPerPageContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageContentType) UnmarshalJSON(b []byte) error {
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

type ApiInfoPerPageStatus struct {
	value int32
}

type ApiInfoPerPageStatusEnum struct {
	E_1 ApiInfoPerPageStatus
}

func GetApiInfoPerPageStatusEnum() ApiInfoPerPageStatusEnum {
	return ApiInfoPerPageStatusEnum{
		E_1: ApiInfoPerPageStatus{
			value: 1,
		},
	}
}

func (c ApiInfoPerPageStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiInfoPerPageStatus) UnmarshalJSON(b []byte) error {
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
