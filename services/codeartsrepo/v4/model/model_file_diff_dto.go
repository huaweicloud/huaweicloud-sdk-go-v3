package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileDiffDto struct {

	// **参数解释：** 旧文件名称。 **取值范围：** 字符串长度不少于1，不超过10000。
	OldPath *string `json:"old_path,omitempty"`

	// **参数解释：** 新文件名称。 **取值范围：** 字符串长度不少于1，不超过10000。
	NewPath *interface{} `json:"new_path,omitempty"`

	// **参数解释：** 旧文件权限。
	AMode *string `json:"a_mode,omitempty"`

	// **参数解释：** 新文件权限。
	BMode *string `json:"b_mode,omitempty"`

	// **参数解释：** 是否为新文件。 **取值范围：** - true，新文件。 - false，非新文件
	NewFile *bool `json:"new_file,omitempty"`

	// **参数解释：** 是否为重命名文件。 **取值范围：** - true，重命名文件。 - false，非重命名文件
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// **参数解释：** 是否为被删除文件。 **取值范围：** - true，被删除文件。 - false，非被删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// **参数解释：** 差异信息。
	Diff *string `json:"diff,omitempty"`

	// **参数解释：** 是否为大文件。 **取值范围：** - true，大文件。 - false，非大文件
	TooLarge *bool `json:"too_large,omitempty"`

	// **参数解释：** 增加行数。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 删除行数。
	RemovedLines *int32 `json:"removed_lines,omitempty"`
}

func (o FileDiffDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileDiffDto struct{}"
	}

	return strings.Join([]string{"FileDiffDto", string(data)}, " ")
}
