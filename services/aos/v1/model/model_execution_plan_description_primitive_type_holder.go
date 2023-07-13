package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlanDescriptionPrimitiveTypeHolder struct {

	// 执行计划的描述。可用于客户识别自己的执行计划。
	Description *string `json:"description,omitempty"`
}

func (o ExecutionPlanDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutionPlanDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
