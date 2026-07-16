package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowStepExecutionPolicy workflow step execution 策略。
type WorkflowStepExecutionPolicy struct {

	// 执行策略，可选值如下： - retry：重试 - stop：停止 - continue：继续运行
	ExecutionPolicy *string `json:"execution_policy,omitempty"`

	// 是否使用的是缓存。
	UseCache *bool `json:"use_cache,omitempty"`
}

func (o WorkflowStepExecutionPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStepExecutionPolicy struct{}"
	}

	return strings.Join([]string{"WorkflowStepExecutionPolicy", string(data)}, " ")
}
