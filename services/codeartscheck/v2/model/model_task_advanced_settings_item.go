package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAdvancedSettingsItem struct {

	// **参数解释：** 高级配置对应的名称。 **约束限制：** 不涉及。 **取值范围：** - customImage：自定义镜像。 - includePaths：检查目录。 **默认取值：** 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 高级配置对应的取值。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~1000]个字符。 **默认取值：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TaskAdvancedSettingsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAdvancedSettingsItem struct{}"
	}

	return strings.Join([]string{"TaskAdvancedSettingsItem", string(data)}, " ")
}
