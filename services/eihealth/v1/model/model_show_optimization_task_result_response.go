package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOptimizationTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *OptimizationTaskData `json:"task_data,omitempty"`

	Result         *OptimizationResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowOptimizationTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOptimizationTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowOptimizationTaskResultResponse", string(data)}, " ")
}
