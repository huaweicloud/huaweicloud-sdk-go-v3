package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PluginTemplateSpecV2 插件模板的具体信息。
type PluginTemplateSpecV2 struct {

	// **参数解释**：是否为必安装插件。 **取值范围**： - true：是 - false：否
	Optional bool `json:"optional"`

	// **参数解释**：插件模板类型。 **取值范围**：可选值如下： - helm：helm类型 - ccePlugin：CCE类型
	Type *PluginTemplateSpecV2Type `json:"type,omitempty"`

	// **参数解释**：Logo图片地址。 **取值范围**：不涉及。
	LogoURL *string `json:"logoURL,omitempty"`

	// **参数解释**：插件模板描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：插件模板版本的详细信息。
	Versions []PluginTemplateVersionV2 `json:"versions"`
}

func (o PluginTemplateSpecV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginTemplateSpecV2 struct{}"
	}

	return strings.Join([]string{"PluginTemplateSpecV2", string(data)}, " ")
}

type PluginTemplateSpecV2Type struct {
	value string
}

type PluginTemplateSpecV2TypeEnum struct {
	HELM       PluginTemplateSpecV2Type
	CCE_PLUGIN PluginTemplateSpecV2Type
}

func GetPluginTemplateSpecV2TypeEnum() PluginTemplateSpecV2TypeEnum {
	return PluginTemplateSpecV2TypeEnum{
		HELM: PluginTemplateSpecV2Type{
			value: "helm",
		},
		CCE_PLUGIN: PluginTemplateSpecV2Type{
			value: " ccePlugin",
		},
	}
}

func (c PluginTemplateSpecV2Type) Value() string {
	return c.value
}

func (c PluginTemplateSpecV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginTemplateSpecV2Type) UnmarshalJSON(b []byte) error {
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
