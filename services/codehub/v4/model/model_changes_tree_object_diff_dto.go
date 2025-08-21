package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangesTreeObjectDiffDto struct {

	// **参数解释：** 变更内容。
	Diff *string `json:"diff,omitempty"`

	// **参数解释：** 变更新路径。
	NewPath *string `json:"new_path,omitempty"`

	// **参数解释：** 变更旧。
	OldPath *string `json:"old_path,omitempty"`

	// **参数解释：** 旧文件权限。
	AMode *string `json:"a_mode,omitempty"`

	// **参数解释：** 新文件权限。
	BMode *string `json:"b_mode,omitempty"`

	// **参数解释：** 是否是新文件。
	NewFile *bool `json:"new_file,omitempty"`

	// **参数解释：** 是否是改名文件。
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// **参数解释：** 是否是删除文件。
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// **参数解释：** 是否内容过长。
	TooLarge *bool `json:"too_large,omitempty"`
}

func (o ChangesTreeObjectDiffDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangesTreeObjectDiffDto struct{}"
	}

	return strings.Join([]string{"ChangesTreeObjectDiffDto", string(data)}, " ")
}
