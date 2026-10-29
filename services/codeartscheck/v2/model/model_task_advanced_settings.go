package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAdvancedSettings struct {

	// **参数解释：** 高级配置对应的名称。 **取值范围：** 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 高级配置对应的取值。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释：** 高级配置对应的可选项。 **取值范围：** 不涉及。
	OptionValue *string `json:"option_value,omitempty"`

	// **参数解释：** 高级配置对应的中文描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o TaskAdvancedSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAdvancedSettings struct{}"
	}

	return strings.Join([]string{"TaskAdvancedSettings", string(data)}, " ")
}
