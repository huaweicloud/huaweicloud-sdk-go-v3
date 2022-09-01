package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiConditionBase struct {

	// 关联的请求参数对象名称。策略类型为param时必选
	ReqParamName *string `json:"req_param_name,omitempty" xml:"req_param_name"`

	// 系统参数名称。策略类型为system时必选。支持以下系统参数 - req_path：请求路径。如 /a/b - req_method：请求方法。如 GET
	SysParamName *ApiConditionBaseSysParamName `json:"sys_param_name,omitempty" xml:"sys_param_name"`

	// 策略条件 - exact：绝对匹配 - enum：枚举 - pattern：正则  策略类型为param时必选
	ConditionType *ApiConditionBaseConditionType `json:"condition_type,omitempty" xml:"condition_type"`

	// 策略类型 - param：参数 - source：源IP - system：系统参数
	ConditionOrigin ApiConditionBaseConditionOrigin `json:"condition_origin" xml:"condition_origin"`

	// 策略值
	ConditionValue string `json:"condition_value" xml:"condition_value"`
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

type ApiConditionBaseConditionOrigin struct {
	value string
}

type ApiConditionBaseConditionOriginEnum struct {
	PARAM  ApiConditionBaseConditionOrigin
	SOURCE ApiConditionBaseConditionOrigin
}

func GetApiConditionBaseConditionOriginEnum() ApiConditionBaseConditionOriginEnum {
	return ApiConditionBaseConditionOriginEnum{
		PARAM: ApiConditionBaseConditionOrigin{
			value: "param",
		},
		SOURCE: ApiConditionBaseConditionOrigin{
			value: "source",
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
