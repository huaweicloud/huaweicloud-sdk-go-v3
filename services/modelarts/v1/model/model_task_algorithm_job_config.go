package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAlgorithmJobConfig **参数解释**：算法配置信息，如启动文件等。 **约束限制**：不涉及。
type TaskAlgorithmJobConfig struct {

	// **参数解释**：算法的运行参数。 **约束限制**：不涉及。
	Parameters *[]Parameter `json:"parameters,omitempty"`

	// **参数解释**：算法的数据输入。 **约束限制**：不涉及。
	Inputs *[]Input `json:"inputs,omitempty"`

	// **参数解释**：算法的数据输出。 **约束限制**：不涉及。
	Outputs *[]Output `json:"outputs,omitempty"`

	Engine *TaskAlgorithmJobConfigEngine `json:"engine,omitempty"`
}

func (o TaskAlgorithmJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmJobConfig struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmJobConfig", string(data)}, " ")
}
