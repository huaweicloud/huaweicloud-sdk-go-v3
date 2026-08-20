package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionExecutionStep 执行步骤
type ExtensionExecutionStep struct {

	// DSL方法名，如 preOperationsNpm/sh/releasemanArtifactsUploader
	DslMethod *string `json:"dslMethod,omitempty"`

	// 步骤显示名
	DisplayName *string `json:"displayName,omitempty"`

	// 执行模式，如 serial
	ExecutionMode *string `json:"executionMode,omitempty"`

	// 步骤参数，键值对，值多为 $${...} 变量引用语法。
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (o ExtensionExecutionStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionExecutionStep struct{}"
	}

	return strings.Join([]string{"ExtensionExecutionStep", string(data)}, " ")
}
