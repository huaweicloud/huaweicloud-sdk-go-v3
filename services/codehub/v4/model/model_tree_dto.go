package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TreeDto struct {

	// **参数解释：** 提交ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 文件名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 文件结构。 **取值范围：** - commit, 子模块。 - tree, 目录。 - blob, 文件
	Type *TreeDtoType `json:"type,omitempty"`

	// **参数解释：** 文件路径。 **取值范围：** 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释：** 当前所在目录层级。 **取值范围：** 不涉及
	Level *int32 `json:"level,omitempty"`

	// **参数解释：** 是否下拉。 **取值范围：** - false, 不下拉。 - true, 下拉
	IsShownDropDown *bool `json:"isShownDropDown,omitempty"`

	// **参数解释：** 是否折叠。 **取值范围：** - false, 不折叠。 - true, 折叠。
	Folder *bool `json:"folder,omitempty"`

	// 子节点
	Children *[]TreeDto `json:"children,omitempty"`

	// **参数解释：** 子模块连接。 **取值范围：** 不涉及。
	SubmoduleLink *string `json:"submodule_link,omitempty"`
}

func (o TreeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeDto struct{}"
	}

	return strings.Join([]string{"TreeDto", string(data)}, " ")
}

type TreeDtoType struct {
	value string
}

type TreeDtoTypeEnum struct {
	COMMIT TreeDtoType
	TREE   TreeDtoType
	BLOB   TreeDtoType
}

func GetTreeDtoTypeEnum() TreeDtoTypeEnum {
	return TreeDtoTypeEnum{
		COMMIT: TreeDtoType{
			value: "commit",
		},
		TREE: TreeDtoType{
			value: "tree",
		},
		BLOB: TreeDtoType{
			value: "blob",
		},
	}
}

func (c TreeDtoType) Value() string {
	return c.value
}

func (c TreeDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TreeDtoType) UnmarshalJSON(b []byte) error {
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
