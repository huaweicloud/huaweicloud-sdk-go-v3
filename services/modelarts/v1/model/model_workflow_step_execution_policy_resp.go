package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowStepExecutionPolicyResp workflow step execution 策略。
type WorkflowStepExecutionPolicyResp struct {

	// **参数解释**：执行策略， **取值范围**：可选值如下： - retry：重试 - stop：停止 - continue：继续运行
	ExecutionPolicy *string `json:"execution_policy,omitempty"`

	// **参数解释**：是否使用的是缓存。 **取值范围**： - true：是缓存 - false：不是缓存
	UseCache *bool `json:"use_cache,omitempty"`
}

func (o WorkflowStepExecutionPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStepExecutionPolicyResp struct{}"
	}

	return strings.Join([]string{"WorkflowStepExecutionPolicyResp", string(data)}, " ")
}
