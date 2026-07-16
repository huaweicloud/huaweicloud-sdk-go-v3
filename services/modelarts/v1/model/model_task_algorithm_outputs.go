package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAlgorithmOutputs struct {

	// **参数解释**：数据输出通道名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：数据输出通道描述信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：数据输出通道映射的容器本地路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LocalDir *string `json:"local_dir,omitempty"`

	Remote *TaskAlgorithmRemote `json:"remote"`
}

func (o TaskAlgorithmOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmOutputs struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmOutputs", string(data)}, " ")
}
