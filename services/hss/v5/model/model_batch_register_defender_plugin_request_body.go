package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchRegisterDefenderPluginRequestBody 批量操作智能体卫士插件请求结构体
type BatchRegisterDefenderPluginRequestBody struct {

	// **参数解释**: 操作类型 **约束限制**: 必填 **取值范围**: - install：安装 - uninstall：卸载  **默认取值**: 不涉及
	Opt BatchRegisterDefenderPluginRequestBodyOpt `json:"opt"`

	// **参数解释**: 智能体实例ID列表 **约束限制**: 必填 **取值范围**: 1-200个实例ID **默认取值**: 不涉及
	InstanceIds []string `json:"instance_ids"`
}

func (o BatchRegisterDefenderPluginRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRegisterDefenderPluginRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRegisterDefenderPluginRequestBody", string(data)}, " ")
}

type BatchRegisterDefenderPluginRequestBodyOpt struct {
	value string
}

type BatchRegisterDefenderPluginRequestBodyOptEnum struct {
	INSTALL   BatchRegisterDefenderPluginRequestBodyOpt
	UNINSTALL BatchRegisterDefenderPluginRequestBodyOpt
}

func GetBatchRegisterDefenderPluginRequestBodyOptEnum() BatchRegisterDefenderPluginRequestBodyOptEnum {
	return BatchRegisterDefenderPluginRequestBodyOptEnum{
		INSTALL: BatchRegisterDefenderPluginRequestBodyOpt{
			value: "install",
		},
		UNINSTALL: BatchRegisterDefenderPluginRequestBodyOpt{
			value: "uninstall",
		},
	}
}

func (c BatchRegisterDefenderPluginRequestBodyOpt) Value() string {
	return c.value
}

func (c BatchRegisterDefenderPluginRequestBodyOpt) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRegisterDefenderPluginRequestBodyOpt) UnmarshalJSON(b []byte) error {
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
