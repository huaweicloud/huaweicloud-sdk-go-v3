package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionExecution 执行定义
type ExtensionExecution struct {

	// 执行步骤列表。
	Steps *[]ExtensionExecutionStep `json:"steps,omitempty"`
}

func (o ExtensionExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionExecution struct{}"
	}

	return strings.Join([]string{"ExtensionExecution", string(data)}, " ")
}
