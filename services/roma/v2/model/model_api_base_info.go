package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiBaseInfo struct {
	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API类型[，该参数暂未使用](tag:hcs;fcs;) - 1：公有API - 2：私有API

	Type ApiBaseInfoType `json:"type"`
	// API的版本

	Version *string `json:"version,omitempty"`
	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS

	ReqProtocol ApiBaseInfoReqProtocol `json:"req_protocol"`
	// API的请求方式

	ReqMethod ApiBaseInfoReqMethod `json:"req_method"`
	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。

	ReqUri string `json:"req_uri"`
	// API的认证方式[，site暂不支持IAM认证。](tag:Site) - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证

	AuthType ApiBaseInfoAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// 是否支持跨域 - TRUE：支持 - FALSE：不支持

	Cors *bool `json:"cors,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL

	MatchMode *ApiBaseInfoMatchMode `json:"match_mode,omitempty"`
	// 后端类型[，site暂不支持函数工作流。](tag:Site) - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端

	BackendType ApiBaseInfoBackendType `json:"backend_type"`
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

	ContentType *ApiBaseInfoContentType `json:"content_type,omitempty"`
}

func (o ApiBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiBaseInfo struct{}"
	}

	return strings.Join([]string{"ApiBaseInfo", string(data)}, " ")
}

type ApiBaseInfoType struct {
	value int32
}

type ApiBaseInfoTypeEnum struct {
	E_1 ApiBaseInfoType
	E_2 ApiBaseInfoType
}

func GetApiBaseInfoTypeEnum() ApiBaseInfoTypeEnum {
	return ApiBaseInfoTypeEnum{
		E_1: ApiBaseInfoType{
			value: 1,
		}, E_2: ApiBaseInfoType{
			value: 2,
		},
	}
}

func (c ApiBaseInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoType) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoReqProtocol struct {
	value string
}

type ApiBaseInfoReqProtocolEnum struct {
	HTTP  ApiBaseInfoReqProtocol
	HTTPS ApiBaseInfoReqProtocol
	BOTH  ApiBaseInfoReqProtocol
}

func GetApiBaseInfoReqProtocolEnum() ApiBaseInfoReqProtocolEnum {
	return ApiBaseInfoReqProtocolEnum{
		HTTP: ApiBaseInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiBaseInfoReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiBaseInfoReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiBaseInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoReqMethod struct {
	value string
}

type ApiBaseInfoReqMethodEnum struct {
	GET     ApiBaseInfoReqMethod
	POST    ApiBaseInfoReqMethod
	PUT     ApiBaseInfoReqMethod
	DELETE  ApiBaseInfoReqMethod
	HEAD    ApiBaseInfoReqMethod
	PATCH   ApiBaseInfoReqMethod
	OPTIONS ApiBaseInfoReqMethod
	ANY     ApiBaseInfoReqMethod
}

func GetApiBaseInfoReqMethodEnum() ApiBaseInfoReqMethodEnum {
	return ApiBaseInfoReqMethodEnum{
		GET: ApiBaseInfoReqMethod{
			value: "GET",
		},
		POST: ApiBaseInfoReqMethod{
			value: "POST",
		},
		PUT: ApiBaseInfoReqMethod{
			value: "PUT",
		},
		DELETE: ApiBaseInfoReqMethod{
			value: "DELETE",
		},
		HEAD: ApiBaseInfoReqMethod{
			value: "HEAD",
		},
		PATCH: ApiBaseInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiBaseInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiBaseInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiBaseInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoAuthType struct {
	value string
}

type ApiBaseInfoAuthTypeEnum struct {
	NONE       ApiBaseInfoAuthType
	APP        ApiBaseInfoAuthType
	IAM        ApiBaseInfoAuthType
	AUTHORIZER ApiBaseInfoAuthType
}

func GetApiBaseInfoAuthTypeEnum() ApiBaseInfoAuthTypeEnum {
	return ApiBaseInfoAuthTypeEnum{
		NONE: ApiBaseInfoAuthType{
			value: "NONE",
		},
		APP: ApiBaseInfoAuthType{
			value: "APP",
		},
		IAM: ApiBaseInfoAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiBaseInfoAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiBaseInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoAuthType) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoMatchMode struct {
	value string
}

type ApiBaseInfoMatchModeEnum struct {
	SWA    ApiBaseInfoMatchMode
	NORMAL ApiBaseInfoMatchMode
}

func GetApiBaseInfoMatchModeEnum() ApiBaseInfoMatchModeEnum {
	return ApiBaseInfoMatchModeEnum{
		SWA: ApiBaseInfoMatchMode{
			value: "SWA",
		},
		NORMAL: ApiBaseInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiBaseInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoBackendType struct {
	value string
}

type ApiBaseInfoBackendTypeEnum struct {
	HTTP     ApiBaseInfoBackendType
	FUNCTION ApiBaseInfoBackendType
	MOCK     ApiBaseInfoBackendType
}

func GetApiBaseInfoBackendTypeEnum() ApiBaseInfoBackendTypeEnum {
	return ApiBaseInfoBackendTypeEnum{
		HTTP: ApiBaseInfoBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiBaseInfoBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiBaseInfoBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiBaseInfoBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoBackendType) UnmarshalJSON(b []byte) error {
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

type ApiBaseInfoContentType struct {
	value string
}

type ApiBaseInfoContentTypeEnum struct {
	APPLICATION_JSON    ApiBaseInfoContentType
	APPLICATION_XML     ApiBaseInfoContentType
	MULTIPART_FORM_DATE ApiBaseInfoContentType
	TEXT_PLAIN          ApiBaseInfoContentType
}

func GetApiBaseInfoContentTypeEnum() ApiBaseInfoContentTypeEnum {
	return ApiBaseInfoContentTypeEnum{
		APPLICATION_JSON: ApiBaseInfoContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiBaseInfoContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiBaseInfoContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiBaseInfoContentType{
			value: "text/plain",
		},
	}
}

func (c ApiBaseInfoContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBaseInfoContentType) UnmarshalJSON(b []byte) error {
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
