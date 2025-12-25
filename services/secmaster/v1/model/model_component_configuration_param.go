package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentConfigurationParam struct {

	// 组件id.
	ComponentId *string `json:"component_id,omitempty"`

	// 配置ID
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	FileType *FileType `json:"file_type,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 参数
	Param *string `json:"param,omitempty"`

	// **参数解释**: 配置类型 - HISTORY 历史版本 - CURRENT_SAVE 当前保存 - CURRENT_APPLY 当前使用  **约束限制** 不涉及 **取值范围**: - HISTORY - CURRENT_SAVE - CURRENT_APPLY  **默认值** 不涉及
	Type *ComponentConfigurationParamType `json:"type,omitempty"`

	// 版本
	Version *int32 `json:"version,omitempty"`
}

func (o ComponentConfigurationParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfigurationParam struct{}"
	}

	return strings.Join([]string{"ComponentConfigurationParam", string(data)}, " ")
}

type ComponentConfigurationParamType struct {
	value string
}

type ComponentConfigurationParamTypeEnum struct {
	HISTORY       ComponentConfigurationParamType
	CURRENT_SAVE  ComponentConfigurationParamType
	CURRENT_APPLY ComponentConfigurationParamType
}

func GetComponentConfigurationParamTypeEnum() ComponentConfigurationParamTypeEnum {
	return ComponentConfigurationParamTypeEnum{
		HISTORY: ComponentConfigurationParamType{
			value: "HISTORY",
		},
		CURRENT_SAVE: ComponentConfigurationParamType{
			value: "CURRENT_SAVE",
		},
		CURRENT_APPLY: ComponentConfigurationParamType{
			value: "CURRENT_APPLY",
		},
	}
}

func (c ComponentConfigurationParamType) Value() string {
	return c.value
}

func (c ComponentConfigurationParamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentConfigurationParamType) UnmarshalJSON(b []byte) error {
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
