package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectSettingsInheritCfgDto struct {

	// **参数解释：** 设置源类型。 **约束限制：** 不涉及。 **取值范围：** - protected_branches，保护分支设置。 - protected_tags，保护Tag设置。 - repository_settings，仓库设置。 - push_rules，提交规则设置。 - merge_requests，合并请求设置。 - e2e_settings，E2E设置。 - watermark，水印设置。 **默认取值：** 不涉及。
	Name *ProjectSettingsInheritCfgDtoName `json:"name,omitempty"`

	// **参数解释：** 继承设置。 **约束限制：** 不涉及。 **取值范围：** - force_inherit，强制继承配置。 - custom，使用当前配置。 **默认取值：** 不涉及。
	InheritMod *ProjectSettingsInheritCfgDtoInheritMod `json:"inherit_mod,omitempty"`
}

func (o ProjectSettingsInheritCfgDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectSettingsInheritCfgDto struct{}"
	}

	return strings.Join([]string{"ProjectSettingsInheritCfgDto", string(data)}, " ")
}

type ProjectSettingsInheritCfgDtoName struct {
	value string
}

type ProjectSettingsInheritCfgDtoNameEnum struct {
	PROTECTED_BRANCHES  ProjectSettingsInheritCfgDtoName
	PROTECTED_TAGS      ProjectSettingsInheritCfgDtoName
	REPOSITORY_SETTINGS ProjectSettingsInheritCfgDtoName
	PUSH_RULES          ProjectSettingsInheritCfgDtoName
	MERGE_REQUESTS      ProjectSettingsInheritCfgDtoName
	E2E_SETTINGS        ProjectSettingsInheritCfgDtoName
	WATERMARK           ProjectSettingsInheritCfgDtoName
}

func GetProjectSettingsInheritCfgDtoNameEnum() ProjectSettingsInheritCfgDtoNameEnum {
	return ProjectSettingsInheritCfgDtoNameEnum{
		PROTECTED_BRANCHES: ProjectSettingsInheritCfgDtoName{
			value: "protected_branches",
		},
		PROTECTED_TAGS: ProjectSettingsInheritCfgDtoName{
			value: "protected_tags",
		},
		REPOSITORY_SETTINGS: ProjectSettingsInheritCfgDtoName{
			value: "repository_settings",
		},
		PUSH_RULES: ProjectSettingsInheritCfgDtoName{
			value: "push_rules",
		},
		MERGE_REQUESTS: ProjectSettingsInheritCfgDtoName{
			value: "merge_requests",
		},
		E2E_SETTINGS: ProjectSettingsInheritCfgDtoName{
			value: "e2e_settings",
		},
		WATERMARK: ProjectSettingsInheritCfgDtoName{
			value: "watermark",
		},
	}
}

func (c ProjectSettingsInheritCfgDtoName) Value() string {
	return c.value
}

func (c ProjectSettingsInheritCfgDtoName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectSettingsInheritCfgDtoName) UnmarshalJSON(b []byte) error {
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

type ProjectSettingsInheritCfgDtoInheritMod struct {
	value string
}

type ProjectSettingsInheritCfgDtoInheritModEnum struct {
	FORCE_INHERIT ProjectSettingsInheritCfgDtoInheritMod
	CUSTOM        ProjectSettingsInheritCfgDtoInheritMod
}

func GetProjectSettingsInheritCfgDtoInheritModEnum() ProjectSettingsInheritCfgDtoInheritModEnum {
	return ProjectSettingsInheritCfgDtoInheritModEnum{
		FORCE_INHERIT: ProjectSettingsInheritCfgDtoInheritMod{
			value: "force_inherit",
		},
		CUSTOM: ProjectSettingsInheritCfgDtoInheritMod{
			value: "custom",
		},
	}
}

func (c ProjectSettingsInheritCfgDtoInheritMod) Value() string {
	return c.value
}

func (c ProjectSettingsInheritCfgDtoInheritMod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectSettingsInheritCfgDtoInheritMod) UnmarshalJSON(b []byte) error {
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
