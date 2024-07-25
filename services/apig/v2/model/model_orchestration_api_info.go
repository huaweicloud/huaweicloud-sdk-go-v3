package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type OrchestrationApiInfo struct {

	// API编号。
	ApiId *string `json:"api_id,omitempty"`

	// API名称。 支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。  > 中文字符必须为UTF-8或者unicode编码。
	ApiName *string `json:"api_name,omitempty"`

	// API的请求方式。
	ReqMethod *OrchestrationApiInfoReqMethod `json:"req_method,omitempty"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ .等特殊字符，总长度不超过512，且满足URI规范。 支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。  > 需要服从URI规范。
	ReqUri *string `json:"req_uri,omitempty"`

	// API的认证方式。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *OrchestrationApiInfoAuthType `json:"auth_type,omitempty"`

	// API的匹配方式。 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *OrchestrationApiInfoMatchMode `json:"match_mode,omitempty"`

	// API所属的分组编号。
	GroupId *string `json:"group_id,omitempty"`

	// API所属分组的名称。
	GroupName *string `json:"group_name,omitempty"`

	// 绑定时间。
	AttachedTime *sdktime.SdkTime `json:"attached_time,omitempty"`
}

func (o OrchestrationApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationApiInfo struct{}"
	}

	return strings.Join([]string{"OrchestrationApiInfo", string(data)}, " ")
}

type OrchestrationApiInfoReqMethod struct {
	value string
}

type OrchestrationApiInfoReqMethodEnum struct {
	GET     OrchestrationApiInfoReqMethod
	POST    OrchestrationApiInfoReqMethod
	PUT     OrchestrationApiInfoReqMethod
	DELETE  OrchestrationApiInfoReqMethod
	HEAD    OrchestrationApiInfoReqMethod
	PATCH   OrchestrationApiInfoReqMethod
	OPTIONS OrchestrationApiInfoReqMethod
	ANY     OrchestrationApiInfoReqMethod
}

func GetOrchestrationApiInfoReqMethodEnum() OrchestrationApiInfoReqMethodEnum {
	return OrchestrationApiInfoReqMethodEnum{
		GET: OrchestrationApiInfoReqMethod{
			value: "GET",
		},
		POST: OrchestrationApiInfoReqMethod{
			value: "POST",
		},
		PUT: OrchestrationApiInfoReqMethod{
			value: "PUT",
		},
		DELETE: OrchestrationApiInfoReqMethod{
			value: "DELETE",
		},
		HEAD: OrchestrationApiInfoReqMethod{
			value: "HEAD",
		},
		PATCH: OrchestrationApiInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: OrchestrationApiInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: OrchestrationApiInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c OrchestrationApiInfoReqMethod) Value() string {
	return c.value
}

func (c OrchestrationApiInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationApiInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type OrchestrationApiInfoAuthType struct {
	value string
}

type OrchestrationApiInfoAuthTypeEnum struct {
	NONE       OrchestrationApiInfoAuthType
	APP        OrchestrationApiInfoAuthType
	IAM        OrchestrationApiInfoAuthType
	AUTHORIZER OrchestrationApiInfoAuthType
}

func GetOrchestrationApiInfoAuthTypeEnum() OrchestrationApiInfoAuthTypeEnum {
	return OrchestrationApiInfoAuthTypeEnum{
		NONE: OrchestrationApiInfoAuthType{
			value: "NONE",
		},
		APP: OrchestrationApiInfoAuthType{
			value: "APP",
		},
		IAM: OrchestrationApiInfoAuthType{
			value: "IAM",
		},
		AUTHORIZER: OrchestrationApiInfoAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c OrchestrationApiInfoAuthType) Value() string {
	return c.value
}

func (c OrchestrationApiInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationApiInfoAuthType) UnmarshalJSON(b []byte) error {
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

type OrchestrationApiInfoMatchMode struct {
	value string
}

type OrchestrationApiInfoMatchModeEnum struct {
	SWA    OrchestrationApiInfoMatchMode
	NORMAL OrchestrationApiInfoMatchMode
}

func GetOrchestrationApiInfoMatchModeEnum() OrchestrationApiInfoMatchModeEnum {
	return OrchestrationApiInfoMatchModeEnum{
		SWA: OrchestrationApiInfoMatchMode{
			value: "SWA",
		},
		NORMAL: OrchestrationApiInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c OrchestrationApiInfoMatchMode) Value() string {
	return c.value
}

func (c OrchestrationApiInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationApiInfoMatchMode) UnmarshalJSON(b []byte) error {
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
