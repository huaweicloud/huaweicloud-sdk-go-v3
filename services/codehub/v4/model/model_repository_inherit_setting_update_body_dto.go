package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryInheritSettingUpdateBodyDto struct {

	// **参数解释：** 设置源类型。 **约束限制：** 不涉及。 **取值范围：** - protected_branches，保护分支设置。 - protected_tags，保护Tag设置。 - repository_settings，仓库设置。 - push_rules，提交规则设置。 - merge_requests，合并请求设置。 - e2e_settings，E2E设置。 - watermark，水印设置。 **默认取值：** 不涉及。
	Name *RepositoryInheritSettingUpdateBodyDtoName `json:"name,omitempty"`

	// **参数解释：** 继承设置。 **约束限制：** 不涉及。 **取值范围：** - inherit，继承上级配置。 - custom，使用当前仓库配置。 **默认取值：** 不涉及。
	InheritMod *RepositoryInheritSettingUpdateBodyDtoInheritMod `json:"inherit_mod,omitempty"`
}

func (o RepositoryInheritSettingUpdateBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryInheritSettingUpdateBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryInheritSettingUpdateBodyDto", string(data)}, " ")
}

type RepositoryInheritSettingUpdateBodyDtoName struct {
	value string
}

type RepositoryInheritSettingUpdateBodyDtoNameEnum struct {
	PROTECTED_BRANCHES  RepositoryInheritSettingUpdateBodyDtoName
	PROTECTED_TAGS      RepositoryInheritSettingUpdateBodyDtoName
	REPOSITORY_SETTINGS RepositoryInheritSettingUpdateBodyDtoName
	PUSH_RULES          RepositoryInheritSettingUpdateBodyDtoName
	MERGE_REQUESTS      RepositoryInheritSettingUpdateBodyDtoName
	E2E_SETTINGS        RepositoryInheritSettingUpdateBodyDtoName
	WATERMARK           RepositoryInheritSettingUpdateBodyDtoName
}

func GetRepositoryInheritSettingUpdateBodyDtoNameEnum() RepositoryInheritSettingUpdateBodyDtoNameEnum {
	return RepositoryInheritSettingUpdateBodyDtoNameEnum{
		PROTECTED_BRANCHES: RepositoryInheritSettingUpdateBodyDtoName{
			value: "protected_branches",
		},
		PROTECTED_TAGS: RepositoryInheritSettingUpdateBodyDtoName{
			value: "protected_tags",
		},
		REPOSITORY_SETTINGS: RepositoryInheritSettingUpdateBodyDtoName{
			value: "repository_settings",
		},
		PUSH_RULES: RepositoryInheritSettingUpdateBodyDtoName{
			value: "push_rules",
		},
		MERGE_REQUESTS: RepositoryInheritSettingUpdateBodyDtoName{
			value: "merge_requests",
		},
		E2E_SETTINGS: RepositoryInheritSettingUpdateBodyDtoName{
			value: "e2e_settings",
		},
		WATERMARK: RepositoryInheritSettingUpdateBodyDtoName{
			value: "watermark",
		},
	}
}

func (c RepositoryInheritSettingUpdateBodyDtoName) Value() string {
	return c.value
}

func (c RepositoryInheritSettingUpdateBodyDtoName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryInheritSettingUpdateBodyDtoName) UnmarshalJSON(b []byte) error {
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

type RepositoryInheritSettingUpdateBodyDtoInheritMod struct {
	value string
}

type RepositoryInheritSettingUpdateBodyDtoInheritModEnum struct {
	INHERIT RepositoryInheritSettingUpdateBodyDtoInheritMod
	CUSTOM  RepositoryInheritSettingUpdateBodyDtoInheritMod
}

func GetRepositoryInheritSettingUpdateBodyDtoInheritModEnum() RepositoryInheritSettingUpdateBodyDtoInheritModEnum {
	return RepositoryInheritSettingUpdateBodyDtoInheritModEnum{
		INHERIT: RepositoryInheritSettingUpdateBodyDtoInheritMod{
			value: "inherit",
		},
		CUSTOM: RepositoryInheritSettingUpdateBodyDtoInheritMod{
			value: "custom",
		},
	}
}

func (c RepositoryInheritSettingUpdateBodyDtoInheritMod) Value() string {
	return c.value
}

func (c RepositoryInheritSettingUpdateBodyDtoInheritMod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryInheritSettingUpdateBodyDtoInheritMod) UnmarshalJSON(b []byte) error {
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
