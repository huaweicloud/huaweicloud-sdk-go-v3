package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ConditionResp struct {
	// 关联的请求参数对象名称。策略类型为param时必选

	ReqParamName *string `json:"req_param_name,omitempty"`
	// 系统参数名称。策略类型为system时必选。支持以下系统参数 - req_path：请求路径。如 /a/b - req_method：请求方法。如 GET

	SysParamName *ConditionRespSysParamName `json:"sys_param_name,omitempty"`
	// 策略条件 - exact：绝对匹配 - enum：枚举 - pattern：正则  策略类型为param时必选

	ConditionType *ConditionRespConditionType `json:"condition_type,omitempty"`
	// 策略类型 - param：参数 - source：源IP - system：系统参数

	ConditionOrigin ConditionRespConditionOrigin `json:"condition_origin"`
	// 策略值

	ConditionValue string `json:"condition_value"`
	// 编号

	Id *string `json:"id,omitempty"`
	// 关联的请求参数对象编号

	ReqParamId *string `json:"req_param_id,omitempty"`
	// 关联的请求参数对象位置

	ReqParamLocation *string `json:"req_param_location,omitempty"`
}

func (o ConditionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionResp struct{}"
	}

	return strings.Join([]string{"ConditionResp", string(data)}, " ")
}

type ConditionRespSysParamName struct {
	value string
}

type ConditionRespSysParamNameEnum struct {
	REQ_PATH   ConditionRespSysParamName
	REQ_METHOD ConditionRespSysParamName
}

func GetConditionRespSysParamNameEnum() ConditionRespSysParamNameEnum {
	return ConditionRespSysParamNameEnum{
		REQ_PATH: ConditionRespSysParamName{
			value: "req_path",
		},
		REQ_METHOD: ConditionRespSysParamName{
			value: "req_method",
		},
	}
}

func (c ConditionRespSysParamName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionRespSysParamName) UnmarshalJSON(b []byte) error {
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

type ConditionRespConditionType struct {
	value string
}

type ConditionRespConditionTypeEnum struct {
	EXACT   ConditionRespConditionType
	ENUM    ConditionRespConditionType
	PATTERN ConditionRespConditionType
}

func GetConditionRespConditionTypeEnum() ConditionRespConditionTypeEnum {
	return ConditionRespConditionTypeEnum{
		EXACT: ConditionRespConditionType{
			value: "exact",
		},
		ENUM: ConditionRespConditionType{
			value: "enum",
		},
		PATTERN: ConditionRespConditionType{
			value: "pattern",
		},
	}
}

func (c ConditionRespConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionRespConditionType) UnmarshalJSON(b []byte) error {
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

type ConditionRespConditionOrigin struct {
	value string
}

type ConditionRespConditionOriginEnum struct {
	PARAM  ConditionRespConditionOrigin
	SOURCE ConditionRespConditionOrigin
}

func GetConditionRespConditionOriginEnum() ConditionRespConditionOriginEnum {
	return ConditionRespConditionOriginEnum{
		PARAM: ConditionRespConditionOrigin{
			value: "param",
		},
		SOURCE: ConditionRespConditionOrigin{
			value: "source",
		},
	}
}

func (c ConditionRespConditionOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionRespConditionOrigin) UnmarshalJSON(b []byte) error {
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
