package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptExecuteParamReference 脚本执行入参引用定义
type ScriptExecuteParamReference struct {

	// 参数引用类型：PARAM_STORE
	ReferType string `json:"refer_type"`

	// 引用参数的唯一主键id
	ParamId ScriptExecuteParamReferenceParamId `json:"param_id"`

	// 引用参数的版本号
	ParamVersion *string `json:"param_version,omitempty"`
}

func (o ScriptExecuteParamReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecuteParamReference struct{}"
	}

	return strings.Join([]string{"ScriptExecuteParamReference", string(data)}, " ")
}

type ScriptExecuteParamReferenceParamId struct {
	value string
}

type ScriptExecuteParamReferenceParamIdEnum struct {
	LOW    ScriptExecuteParamReferenceParamId
	MEDIUM ScriptExecuteParamReferenceParamId
	HIGH   ScriptExecuteParamReferenceParamId
}

func GetScriptExecuteParamReferenceParamIdEnum() ScriptExecuteParamReferenceParamIdEnum {
	return ScriptExecuteParamReferenceParamIdEnum{
		LOW: ScriptExecuteParamReferenceParamId{
			value: "LOW",
		},
		MEDIUM: ScriptExecuteParamReferenceParamId{
			value: "MEDIUM",
		},
		HIGH: ScriptExecuteParamReferenceParamId{
			value: "HIGH",
		},
	}
}

func (c ScriptExecuteParamReferenceParamId) Value() string {
	return c.value
}

func (c ScriptExecuteParamReferenceParamId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptExecuteParamReferenceParamId) UnmarshalJSON(b []byte) error {
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
