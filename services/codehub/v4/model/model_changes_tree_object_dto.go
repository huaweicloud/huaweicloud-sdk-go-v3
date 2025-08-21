package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangesTreeObjectDto struct {

	// **参数解释：** 变更文件名。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 文件层级。
	Level *int32 `json:"level,omitempty"`

	// **参数解释：** 文件路径。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 文件类型。
	FileType *string `json:"file_type,omitempty"`

	Diff *ChangesTreeObjectDiffDto `json:"diff,omitempty"`

	// **参数解释：** 子文件变更。
	Items *[]ChangesTreeObjectDto `json:"items,omitempty"`
}

func (o ChangesTreeObjectDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangesTreeObjectDto struct{}"
	}

	return strings.Join([]string{"ChangesTreeObjectDto", string(data)}, " ")
}
