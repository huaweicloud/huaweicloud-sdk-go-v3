package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryInheritSettingDto struct {

	// **参数解释：** 设置源类型。 **取值范围：** - protected_branches，保护分支设置。 - protected_tags，保护Tag设置。 - repository_settings，仓库设置。 - push_rules，提交规则设置。 - merge_requests，合并请求设置。 - e2e_settings，E2E设置。 - watermark，水印设置。
	Name *RepositoryInheritSettingDtoName `json:"name,omitempty"`

	// **参数解释：** 继承设置。 **取值范围：** - inherit，继承上级配置。 - custom，使用当前仓库配置。 - force_inherit，强制继承上级配置。
	InheritMod *RepositoryInheritSettingDtoInheritMod `json:"inherit_mod,omitempty"`
}

func (o RepositoryInheritSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryInheritSettingDto struct{}"
	}

	return strings.Join([]string{"RepositoryInheritSettingDto", string(data)}, " ")
}

type RepositoryInheritSettingDtoName struct {
	value string
}

type RepositoryInheritSettingDtoNameEnum struct {
	PROTECTED_BRANCHES  RepositoryInheritSettingDtoName
	PROTECTED_TAGS      RepositoryInheritSettingDtoName
	REPOSITORY_SETTINGS RepositoryInheritSettingDtoName
	PUSH_RULES          RepositoryInheritSettingDtoName
	MERGE_REQUESTS      RepositoryInheritSettingDtoName
	E2E_SETTINGS        RepositoryInheritSettingDtoName
	WATERMARK           RepositoryInheritSettingDtoName
}

func GetRepositoryInheritSettingDtoNameEnum() RepositoryInheritSettingDtoNameEnum {
	return RepositoryInheritSettingDtoNameEnum{
		PROTECTED_BRANCHES: RepositoryInheritSettingDtoName{
			value: "protected_branches",
		},
		PROTECTED_TAGS: RepositoryInheritSettingDtoName{
			value: "protected_tags",
		},
		REPOSITORY_SETTINGS: RepositoryInheritSettingDtoName{
			value: "repository_settings",
		},
		PUSH_RULES: RepositoryInheritSettingDtoName{
			value: "push_rules",
		},
		MERGE_REQUESTS: RepositoryInheritSettingDtoName{
			value: "merge_requests",
		},
		E2E_SETTINGS: RepositoryInheritSettingDtoName{
			value: "e2e_settings",
		},
		WATERMARK: RepositoryInheritSettingDtoName{
			value: "watermark",
		},
	}
}

func (c RepositoryInheritSettingDtoName) Value() string {
	return c.value
}

func (c RepositoryInheritSettingDtoName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryInheritSettingDtoName) UnmarshalJSON(b []byte) error {
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

type RepositoryInheritSettingDtoInheritMod struct {
	value string
}

type RepositoryInheritSettingDtoInheritModEnum struct {
	INHERIT       RepositoryInheritSettingDtoInheritMod
	CUSTOM        RepositoryInheritSettingDtoInheritMod
	FORCE_INHERIT RepositoryInheritSettingDtoInheritMod
}

func GetRepositoryInheritSettingDtoInheritModEnum() RepositoryInheritSettingDtoInheritModEnum {
	return RepositoryInheritSettingDtoInheritModEnum{
		INHERIT: RepositoryInheritSettingDtoInheritMod{
			value: "inherit",
		},
		CUSTOM: RepositoryInheritSettingDtoInheritMod{
			value: "custom",
		},
		FORCE_INHERIT: RepositoryInheritSettingDtoInheritMod{
			value: "force_inherit",
		},
	}
}

func (c RepositoryInheritSettingDtoInheritMod) Value() string {
	return c.value
}

func (c RepositoryInheritSettingDtoInheritMod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryInheritSettingDtoInheritMod) UnmarshalJSON(b []byte) error {
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
