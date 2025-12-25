package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FileConfiguration struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	FileType *FileType `json:"file_type"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 参数
	Param map[string]string `json:"param"`

	// **参数解释**: 配置类型 - HISTORY 历史版本 - CURRENT_SAVE 当前保存 - CURRENT_APPLY 当前使用  **约束限制** 不涉及 **取值范围**: - HISTORY - CURRENT_SAVE - CURRENT_APPLY  **默认值** 不涉及
	Type *FileConfigurationType `json:"type,omitempty"`
}

func (o FileConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileConfiguration struct{}"
	}

	return strings.Join([]string{"FileConfiguration", string(data)}, " ")
}

type FileConfigurationType struct {
	value string
}

type FileConfigurationTypeEnum struct {
	HISTORY       FileConfigurationType
	CURRENT_SAVE  FileConfigurationType
	CURRENT_APPLY FileConfigurationType
}

func GetFileConfigurationTypeEnum() FileConfigurationTypeEnum {
	return FileConfigurationTypeEnum{
		HISTORY: FileConfigurationType{
			value: "HISTORY",
		},
		CURRENT_SAVE: FileConfigurationType{
			value: "CURRENT_SAVE",
		},
		CURRENT_APPLY: FileConfigurationType{
			value: "CURRENT_APPLY",
		},
	}
}

func (c FileConfigurationType) Value() string {
	return c.value
}

func (c FileConfigurationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileConfigurationType) UnmarshalJSON(b []byte) error {
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
