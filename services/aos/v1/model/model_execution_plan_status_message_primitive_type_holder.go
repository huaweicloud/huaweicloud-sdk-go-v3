package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlanStatusMessagePrimitiveTypeHolder struct {

	// 当执行计划的状态为创建失败状态（即为 `CREATION_FAILED` 时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o ExecutionPlanStatusMessagePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanStatusMessagePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutionPlanStatusMessagePrimitiveTypeHolder", string(data)}, " ")
}
