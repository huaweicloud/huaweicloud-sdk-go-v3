package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiConditionBase struct {

	// 关联的请求参数对象名称。策略类型为param时必选
	ReqParamName *string `json:"req_param_name,omitempty"`

	// 系统参数-网关内置参数名称。策略类型为system时必选。支持以下参数 - req_path：请求路径。如 /a/b - req_method：请求方法。如 GET
	SysParamName *ApiConditionBaseSysParamName `json:"sys_param_name,omitempty"`

	// COOKIE参数名称。策略类型为cookie时必选
	CookieParamName *string `json:"cookie_param_name,omitempty"`

	// 系统参数-前端认证参数名称。策略类型为frontend_authorizer时必选，前端认证参数名称以\"$context.authorizer.frontend.\"字符串为前缀。例如，前端认证参数名称为user_name，加上前缀为$context.authorizer.frontend.user_name。
	FrontendAuthorizerParamName *string `json:"frontend_authorizer_param_name,omitempty"`

	// 策略条件 - exact：绝对匹配 - enum：枚举 - pattern：正则  策略类型为param，system，cookie，frontend_authorizer时必选
	ConditionType *ApiConditionBaseConditionType `json:"condition_type,omitempty"`

	// 策略类型 - param：参数 - source：源IP - system: 系统参数-网关内置参数 - cookie: COOKIE参数 - frontend_authorizer: 系统参数-前端认证参数
	ConditionOrigin ApiConditionBaseConditionOrigin `json:"condition_origin"`

	// 策略值。
	ConditionValue string `json:"condition_value"`

	// 参数编排规则编排后生成的参数名称，当condition_origin为orchestration的时候必填，并且生成的参数名称必须在api绑定的编排规则中存在
	MappedParamName *string `json:"mapped_param_name,omitempty"`

	// 参数编排规则编排后生成的参数所在的位置，当condition_origin为orchestration的时候必填，并且生成的参数所在的位置必须在api绑定的编排规则中存在
	MappedParamLocation *ApiConditionBaseMappedParamLocation `json:"mapped_param_location,omitempty"`
}

func (o ApiConditionBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiConditionBase struct{}"
	}

	return strings.Join([]string{"ApiConditionBase", string(data)}, " ")
}

type ApiConditionBaseSysParamName struct {
	value string
}

type ApiConditionBaseSysParamNameEnum struct {
	REQ_PATH   ApiConditionBaseSysParamName
	REQ_METHOD ApiConditionBaseSysParamName
}

func GetApiConditionBaseSysParamNameEnum() ApiConditionBaseSysParamNameEnum {
	return ApiConditionBaseSysParamNameEnum{
		REQ_PATH: ApiConditionBaseSysParamName{
			value: "req_path",
		},
		REQ_METHOD: ApiConditionBaseSysParamName{
			value: "req_method",
		},
	}
}

func (c ApiConditionBaseSysParamName) Value() string {
	return c.value
}

func (c ApiConditionBaseSysParamName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionBaseSysParamName) UnmarshalJSON(b []byte) error {
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

type ApiConditionBaseConditionType struct {
	value string
}

type ApiConditionBaseConditionTypeEnum struct {
	EXACT   ApiConditionBaseConditionType
	ENUM    ApiConditionBaseConditionType
	PATTERN ApiConditionBaseConditionType
}

func GetApiConditionBaseConditionTypeEnum() ApiConditionBaseConditionTypeEnum {
	return ApiConditionBaseConditionTypeEnum{
		EXACT: ApiConditionBaseConditionType{
			value: "exact",
		},
		ENUM: ApiConditionBaseConditionType{
			value: "enum",
		},
		PATTERN: ApiConditionBaseConditionType{
			value: "pattern",
		},
	}
}

func (c ApiConditionBaseConditionType) Value() string {
	return c.value
}

func (c ApiConditionBaseConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionBaseConditionType) UnmarshalJSON(b []byte) error {
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

type ApiConditionBaseConditionOrigin struct {
	value string
}

type ApiConditionBaseConditionOriginEnum struct {
	PARAM               ApiConditionBaseConditionOrigin
	SOURCE              ApiConditionBaseConditionOrigin
	SYSTEM              ApiConditionBaseConditionOrigin
	COOKIE              ApiConditionBaseConditionOrigin
	FRONTEND_AUTHORIZER ApiConditionBaseConditionOrigin
}

func GetApiConditionBaseConditionOriginEnum() ApiConditionBaseConditionOriginEnum {
	return ApiConditionBaseConditionOriginEnum{
		PARAM: ApiConditionBaseConditionOrigin{
			value: "param",
		},
		SOURCE: ApiConditionBaseConditionOrigin{
			value: "source",
		},
		SYSTEM: ApiConditionBaseConditionOrigin{
			value: "system",
		},
		COOKIE: ApiConditionBaseConditionOrigin{
			value: "cookie",
		},
		FRONTEND_AUTHORIZER: ApiConditionBaseConditionOrigin{
			value: "frontend_authorizer",
		},
	}
}

func (c ApiConditionBaseConditionOrigin) Value() string {
	return c.value
}

func (c ApiConditionBaseConditionOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionBaseConditionOrigin) UnmarshalJSON(b []byte) error {
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

type ApiConditionBaseMappedParamLocation struct {
	value string
}

type ApiConditionBaseMappedParamLocationEnum struct {
	HEADER ApiConditionBaseMappedParamLocation
	QUERY  ApiConditionBaseMappedParamLocation
}

func GetApiConditionBaseMappedParamLocationEnum() ApiConditionBaseMappedParamLocationEnum {
	return ApiConditionBaseMappedParamLocationEnum{
		HEADER: ApiConditionBaseMappedParamLocation{
			value: "header",
		},
		QUERY: ApiConditionBaseMappedParamLocation{
			value: "query",
		},
	}
}

func (c ApiConditionBaseMappedParamLocation) Value() string {
	return c.value
}

func (c ApiConditionBaseMappedParamLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionBaseMappedParamLocation) UnmarshalJSON(b []byte) error {
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
