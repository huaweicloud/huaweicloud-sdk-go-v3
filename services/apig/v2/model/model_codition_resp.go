package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CoditionResp struct {

	// 关联的请求参数对象名称。策略类型为param时必选
	ReqParamName *string `json:"req_param_name,omitempty"`

	// 系统参数-网关内置参数名称。策略类型为system时必选。支持以下参数 - req_path：请求路径。如 /a/b - req_method：请求方法。如 GET
	SysParamName *CoditionRespSysParamName `json:"sys_param_name,omitempty"`

	// COOKIE参数名称。策略类型为cookie时必选
	CookieParamName *string `json:"cookie_param_name,omitempty"`

	// 系统参数-前端认证参数名称。策略类型为frontend_authorizer时必选，前端认证参数名称以\"$context.authorizer.frontend.\"字符串为前缀。例如，前端认证参数名称为user_name，加上前缀为$context.authorizer.frontend.user_name。
	FrontendAuthorizerParamName *string `json:"frontend_authorizer_param_name,omitempty"`

	// 策略条件 - exact：绝对匹配 - enum：枚举 - pattern：正则  策略类型为param，system，cookie，frontend_authorizer时必选
	ConditionType *CoditionRespConditionType `json:"condition_type,omitempty"`

	// 策略类型 - param：参数 - source：源IP - system: 系统参数-网关内置参数 - cookie: COOKIE参数 - frontend_authorizer: 系统参数-前端认证参数
	ConditionOrigin CoditionRespConditionOrigin `json:"condition_origin"`

	// 策略值。
	ConditionValue string `json:"condition_value"`

	// 参数编排规则编排后生成的参数名称，当condition_origin为orchestration的时候必填，并且生成的参数名称必须在api绑定的编排规则中存在
	MappedParamName *string `json:"mapped_param_name,omitempty"`

	// 参数编排规则编排后生成的参数所在的位置，当condition_origin为orchestration的时候必填，并且生成的参数所在的位置必须在api绑定的编排规则中存在
	MappedParamLocation *CoditionRespMappedParamLocation `json:"mapped_param_location,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 关联的请求参数对象编号
	ReqParamId *string `json:"req_param_id,omitempty"`

	// 关联的请求参数对象位置
	ReqParamLocation *string `json:"req_param_location,omitempty"`
}

func (o CoditionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoditionResp struct{}"
	}

	return strings.Join([]string{"CoditionResp", string(data)}, " ")
}

type CoditionRespSysParamName struct {
	value string
}

type CoditionRespSysParamNameEnum struct {
	REQ_PATH   CoditionRespSysParamName
	REQ_METHOD CoditionRespSysParamName
}

func GetCoditionRespSysParamNameEnum() CoditionRespSysParamNameEnum {
	return CoditionRespSysParamNameEnum{
		REQ_PATH: CoditionRespSysParamName{
			value: "req_path",
		},
		REQ_METHOD: CoditionRespSysParamName{
			value: "req_method",
		},
	}
}

func (c CoditionRespSysParamName) Value() string {
	return c.value
}

func (c CoditionRespSysParamName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoditionRespSysParamName) UnmarshalJSON(b []byte) error {
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

type CoditionRespConditionType struct {
	value string
}

type CoditionRespConditionTypeEnum struct {
	EXACT   CoditionRespConditionType
	ENUM    CoditionRespConditionType
	PATTERN CoditionRespConditionType
}

func GetCoditionRespConditionTypeEnum() CoditionRespConditionTypeEnum {
	return CoditionRespConditionTypeEnum{
		EXACT: CoditionRespConditionType{
			value: "exact",
		},
		ENUM: CoditionRespConditionType{
			value: "enum",
		},
		PATTERN: CoditionRespConditionType{
			value: "pattern",
		},
	}
}

func (c CoditionRespConditionType) Value() string {
	return c.value
}

func (c CoditionRespConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoditionRespConditionType) UnmarshalJSON(b []byte) error {
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

type CoditionRespConditionOrigin struct {
	value string
}

type CoditionRespConditionOriginEnum struct {
	PARAM               CoditionRespConditionOrigin
	SOURCE              CoditionRespConditionOrigin
	SYSTEM              CoditionRespConditionOrigin
	COOKIE              CoditionRespConditionOrigin
	FRONTEND_AUTHORIZER CoditionRespConditionOrigin
}

func GetCoditionRespConditionOriginEnum() CoditionRespConditionOriginEnum {
	return CoditionRespConditionOriginEnum{
		PARAM: CoditionRespConditionOrigin{
			value: "param",
		},
		SOURCE: CoditionRespConditionOrigin{
			value: "source",
		},
		SYSTEM: CoditionRespConditionOrigin{
			value: "system",
		},
		COOKIE: CoditionRespConditionOrigin{
			value: "cookie",
		},
		FRONTEND_AUTHORIZER: CoditionRespConditionOrigin{
			value: "frontend_authorizer",
		},
	}
}

func (c CoditionRespConditionOrigin) Value() string {
	return c.value
}

func (c CoditionRespConditionOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoditionRespConditionOrigin) UnmarshalJSON(b []byte) error {
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

type CoditionRespMappedParamLocation struct {
	value string
}

type CoditionRespMappedParamLocationEnum struct {
	HEADER CoditionRespMappedParamLocation
	QUERY  CoditionRespMappedParamLocation
}

func GetCoditionRespMappedParamLocationEnum() CoditionRespMappedParamLocationEnum {
	return CoditionRespMappedParamLocationEnum{
		HEADER: CoditionRespMappedParamLocation{
			value: "header",
		},
		QUERY: CoditionRespMappedParamLocation{
			value: "query",
		},
	}
}

func (c CoditionRespMappedParamLocation) Value() string {
	return c.value
}

func (c CoditionRespMappedParamLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoditionRespMappedParamLocation) UnmarshalJSON(b []byte) error {
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
