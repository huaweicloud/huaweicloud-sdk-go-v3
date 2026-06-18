package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ManageableGroupDto struct {

	// **参数解释：** 代码组全名。
	FullName *string `json:"full_name,omitempty"`

	// **参数解释：** 代码组id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 代码组名。 **取值范围：** 字符串长度不少于0，不超过256。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 全路径。 **取值范围：** 字符串长度不少于0，不超过1000。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释：** 路径。 **取值范围：** 字符串长度不少于0，不超过1000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 可见性。 **取值范围：** private public。
	Visibility *ManageableGroupDtoVisibility `json:"visibility,omitempty"`
}

func (o ManageableGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManageableGroupDto struct{}"
	}

	return strings.Join([]string{"ManageableGroupDto", string(data)}, " ")
}

type ManageableGroupDtoVisibility struct {
	value string
}

type ManageableGroupDtoVisibilityEnum struct {
	PUBLIC  ManageableGroupDtoVisibility
	PRIVATE ManageableGroupDtoVisibility
}

func GetManageableGroupDtoVisibilityEnum() ManageableGroupDtoVisibilityEnum {
	return ManageableGroupDtoVisibilityEnum{
		PUBLIC: ManageableGroupDtoVisibility{
			value: "public",
		},
		PRIVATE: ManageableGroupDtoVisibility{
			value: "private",
		},
	}
}

func (c ManageableGroupDtoVisibility) Value() string {
	return c.value
}

func (c ManageableGroupDtoVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManageableGroupDtoVisibility) UnmarshalJSON(b []byte) error {
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
