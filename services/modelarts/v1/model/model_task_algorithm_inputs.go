package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAlgorithmInputs struct {

	// **参数解释**：数据输入通道名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：数据输入通道描述信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：数据输入通道映射的容器本地路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LocalDir *string `json:"local_dir,omitempty"`

	Remote *InputDataInfo `json:"remote"`
}

func (o TaskAlgorithmInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmInputs struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmInputs", string(data)}, " ")
}
