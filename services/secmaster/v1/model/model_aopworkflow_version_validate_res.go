package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AopworkflowVersionValidateRes 校验流程的请求对象
type AopworkflowVersionValidateRes struct {

	// **参数解释**: 流程的校验类型 - BASIC 基础校验 - CIRCLE 环路校验 - APP_PARAMS 参数校验  **约束限制**: 不涉及         **取值范围**: - BASIC - CIRCLE - APP_PARAMS  **默认值**:  不涉及
	Mode AopworkflowVersionValidateResMode `json:"mode"`

	// **参数解释**: 流程拓扑图的参数配置 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskconfig string `json:"taskconfig"`

	// **参数解释**: 流程的拓扑图的BASE64编码 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskflow string `json:"taskflow"`

	// **参数解释**: 流程的ID **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	AopworkflowId string `json:"aopworkflow_id"`
}

func (o AopworkflowVersionValidateRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopworkflowVersionValidateRes struct{}"
	}

	return strings.Join([]string{"AopworkflowVersionValidateRes", string(data)}, " ")
}

type AopworkflowVersionValidateResMode struct {
	value string
}

type AopworkflowVersionValidateResModeEnum struct {
	BASIC      AopworkflowVersionValidateResMode
	CIRCLE     AopworkflowVersionValidateResMode
	APP_PARAMS AopworkflowVersionValidateResMode
}

func GetAopworkflowVersionValidateResModeEnum() AopworkflowVersionValidateResModeEnum {
	return AopworkflowVersionValidateResModeEnum{
		BASIC: AopworkflowVersionValidateResMode{
			value: "BASIC",
		},
		CIRCLE: AopworkflowVersionValidateResMode{
			value: "CIRCLE",
		},
		APP_PARAMS: AopworkflowVersionValidateResMode{
			value: "APP_PARAMS",
		},
	}
}

func (c AopworkflowVersionValidateResMode) Value() string {
	return c.value
}

func (c AopworkflowVersionValidateResMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AopworkflowVersionValidateResMode) UnmarshalJSON(b []byte) error {
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
