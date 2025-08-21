package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResolveConflictFileDto struct {

	// **参数解释：** 旧文件路径
	OldPath string `json:"old_path"`

	// **参数解释：** 新文件路径
	NewPath string `json:"new_path"`

	// **参数解释：** 只有选择接受不同分支选项的时候才需要
	Sections *interface{} `json:"sections,omitempty"`

	// **参数解释：** 只有在线编辑冲突文件的时候才需要，内容即为文件内容
	Content *string `json:"content,omitempty"`
}

func (o ResolveConflictFileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolveConflictFileDto struct{}"
	}

	return strings.Join([]string{"ResolveConflictFileDto", string(data)}, " ")
}
