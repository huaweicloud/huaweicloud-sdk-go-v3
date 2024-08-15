package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScriptExecuteInputParam struct {

	// 脚本入参的名称,同一个脚本，参数名不能重复
	ParamName string `json:"param_name"`

	// 脚本入参的值，默认必填。有引用参数（param_refer不为空）时，允许为空
	ParamValue ScriptExecuteInputParamParamValue `json:"param_value"`

	// 该参数已废弃，传入该参数不会生效。
	ParamOrder *int32 `json:"param_order,omitempty"`

	ParamRefer *ScriptExecuteParamReference `json:"param_refer,omitempty"`
}

func (o ScriptExecuteInputParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecuteInputParam struct{}"
	}

	return strings.Join([]string{"ScriptExecuteInputParam", string(data)}, " ")
}

type ScriptExecuteInputParamParamValue struct {
	value string
}

type ScriptExecuteInputParamParamValueEnum struct {
	___2                  ScriptExecuteInputParamParamValue
	A_Z_A_Z0_9________X20 ScriptExecuteInputParamParamValue
	_______               ScriptExecuteInputParamParamValue
}

func GetScriptExecuteInputParamParamValueEnum() ScriptExecuteInputParamParamValueEnum {
	return ScriptExecuteInputParamParamValueEnum{
		___2: ScriptExecuteInputParamParamValue{
			value: "^((?!\\.{2",
		},
		A_Z_A_Z0_9________X20: ScriptExecuteInputParamParamValue{
			value: "})[a-zA-Z0-9_\\-/\\.\\x20\\?:\"",
		},
		_______: ScriptExecuteInputParamParamValue{
			value: "=+@\\\\\\[\\{\\]\\}])*$",
		},
	}
}

func (c ScriptExecuteInputParamParamValue) Value() string {
	return c.value
}

func (c ScriptExecuteInputParamParamValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptExecuteInputParamParamValue) UnmarshalJSON(b []byte) error {
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
