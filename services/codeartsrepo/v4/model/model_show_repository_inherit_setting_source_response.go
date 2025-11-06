package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryInheritSettingSourceResponse Response Object
type ShowRepositoryInheritSettingSourceResponse struct {

	// **参数解释：** 设置源类型。 **取值范围：** - project，项目。 - group，代码组。 - repository，仓库。
	SourceType *ShowRepositoryInheritSettingSourceResponseSourceType `json:"source_type,omitempty"`

	// **参数解释：** 设置源ID。 **取值范围：** 1-2147483647
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释：** 继承设置是否可修改。 **取值范围：** - true，可修改。 - false，不可修改。
	UpwardInheritEditable *bool `json:"upward_inherit_editable,omitempty"`
	HttpStatusCode        int   `json:"-"`
}

func (o ShowRepositoryInheritSettingSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryInheritSettingSourceResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryInheritSettingSourceResponse", string(data)}, " ")
}

type ShowRepositoryInheritSettingSourceResponseSourceType struct {
	value string
}

type ShowRepositoryInheritSettingSourceResponseSourceTypeEnum struct {
	PROJECT    ShowRepositoryInheritSettingSourceResponseSourceType
	GROUP      ShowRepositoryInheritSettingSourceResponseSourceType
	REPOSITORY ShowRepositoryInheritSettingSourceResponseSourceType
}

func GetShowRepositoryInheritSettingSourceResponseSourceTypeEnum() ShowRepositoryInheritSettingSourceResponseSourceTypeEnum {
	return ShowRepositoryInheritSettingSourceResponseSourceTypeEnum{
		PROJECT: ShowRepositoryInheritSettingSourceResponseSourceType{
			value: "project",
		},
		GROUP: ShowRepositoryInheritSettingSourceResponseSourceType{
			value: "group",
		},
		REPOSITORY: ShowRepositoryInheritSettingSourceResponseSourceType{
			value: "repository",
		},
	}
}

func (c ShowRepositoryInheritSettingSourceResponseSourceType) Value() string {
	return c.value
}

func (c ShowRepositoryInheritSettingSourceResponseSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryInheritSettingSourceResponseSourceType) UnmarshalJSON(b []byte) error {
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
