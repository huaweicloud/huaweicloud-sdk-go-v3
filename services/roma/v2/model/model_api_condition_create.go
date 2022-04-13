package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiConditionCreate struct {
	// 关联的请求参数对象名称。策略类型为param时必选

	ReqParamName *string `json:"req_param_name,omitempty"`
	// 系统参数名称。策略类型为system时必选。支持以下系统参数 - req_path：请求路径。如 /a/b - req_method：请求方法。如 GET

	SysParamName *ApiConditionCreateSysParamName `json:"sys_param_name,omitempty"`
	// 策略条件 - exact：绝对匹配 - enum：枚举 - pattern：正则  策略类型为param时必选

	ConditionType *ApiConditionCreateConditionType `json:"condition_type,omitempty"`
	// 策略类型 - param：参数 - source：源IP - system：系统参数

	ConditionOrigin ApiConditionCreateConditionOrigin `json:"condition_origin"`
	// 策略值

	ConditionValue string `json:"condition_value"`
}

func (o ApiConditionCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiConditionCreate struct{}"
	}

	return strings.Join([]string{"ApiConditionCreate", string(data)}, " ")
}

type ApiConditionCreateSysParamName struct {
	value string
}

type ApiConditionCreateSysParamNameEnum struct {
	REQ_PATH   ApiConditionCreateSysParamName
	REQ_METHOD ApiConditionCreateSysParamName
}

func GetApiConditionCreateSysParamNameEnum() ApiConditionCreateSysParamNameEnum {
	return ApiConditionCreateSysParamNameEnum{
		REQ_PATH: ApiConditionCreateSysParamName{
			value: "req_path",
		},
		REQ_METHOD: ApiConditionCreateSysParamName{
			value: "req_method",
		},
	}
}

func (c ApiConditionCreateSysParamName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionCreateSysParamName) UnmarshalJSON(b []byte) error {
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

type ApiConditionCreateConditionType struct {
	value string
}

type ApiConditionCreateConditionTypeEnum struct {
	EXACT   ApiConditionCreateConditionType
	ENUM    ApiConditionCreateConditionType
	PATTERN ApiConditionCreateConditionType
}

func GetApiConditionCreateConditionTypeEnum() ApiConditionCreateConditionTypeEnum {
	return ApiConditionCreateConditionTypeEnum{
		EXACT: ApiConditionCreateConditionType{
			value: "exact",
		},
		ENUM: ApiConditionCreateConditionType{
			value: "enum",
		},
		PATTERN: ApiConditionCreateConditionType{
			value: "pattern",
		},
	}
}

func (c ApiConditionCreateConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionCreateConditionType) UnmarshalJSON(b []byte) error {
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

type ApiConditionCreateConditionOrigin struct {
	value string
}

type ApiConditionCreateConditionOriginEnum struct {
	PARAM  ApiConditionCreateConditionOrigin
	SOURCE ApiConditionCreateConditionOrigin
}

func GetApiConditionCreateConditionOriginEnum() ApiConditionCreateConditionOriginEnum {
	return ApiConditionCreateConditionOriginEnum{
		PARAM: ApiConditionCreateConditionOrigin{
			value: "param",
		},
		SOURCE: ApiConditionCreateConditionOrigin{
			value: "source",
		},
	}
}

func (c ApiConditionCreateConditionOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiConditionCreateConditionOrigin) UnmarshalJSON(b []byte) error {
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
