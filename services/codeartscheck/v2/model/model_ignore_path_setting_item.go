package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IgnorePathSettingItem struct {

	// **参数解释：** 需要屏蔽的目录或文件路径。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 屏蔽状态。 **约束限制：** 不涉及。 **取值范围：** - unchecked：取消屏蔽 - all：全部屏蔽 - half：半屏蔽，用于目录，表示目录下部分子目录、文件需要屏蔽 **默认取值：** 不涉及。
	CheckboxStatus *string `json:"checkbox_status,omitempty"`
}

func (o IgnorePathSettingItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnorePathSettingItem struct{}"
	}

	return strings.Join([]string{"IgnorePathSettingItem", string(data)}, " ")
}
