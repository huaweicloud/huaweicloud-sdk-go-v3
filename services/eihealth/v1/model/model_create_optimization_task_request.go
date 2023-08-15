package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOptimizationTaskRequest Request Object
type CreateOptimizationTaskRequest struct {
	Body *OptimizationTaskData `json:"body,omitempty"`
}

func (o CreateOptimizationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOptimizationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateOptimizationTaskRequest", string(data)}, " ")
}
