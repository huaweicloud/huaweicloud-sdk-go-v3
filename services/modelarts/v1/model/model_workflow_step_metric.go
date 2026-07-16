package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowStepMetric 工作流节点度量信息。
type WorkflowStepMetric struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 度量项。
	Key *string `json:"key,omitempty"`

	// 度量标题。
	Title *string `json:"title,omitempty"`

	// 度量的类型。
	Type *string `json:"type,omitempty"`

	// 度量数据。
	Data map[string]interface{} `json:"data,omitempty"`
}

func (o WorkflowStepMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStepMetric struct{}"
	}

	return strings.Join([]string{"WorkflowStepMetric", string(data)}, " ")
}
