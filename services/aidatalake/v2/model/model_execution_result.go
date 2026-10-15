package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutionResult **参数解释**：作业执行结果。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
type ExecutionResult struct {

	// **参数解释**：作业执行退出码。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	ExitCode *string `json:"exit_code,omitempty"`

	// **参数解释**：作业执行错误信息。 **约束限制**：不涉及。 **取值范围**：长度为1~1024的任意字符组合。 **默认取值**：不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ExecutionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionResult struct{}"
	}

	return strings.Join([]string{"ExecutionResult", string(data)}, " ")
}
