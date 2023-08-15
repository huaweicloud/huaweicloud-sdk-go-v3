package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOptimizationTaskResultRequest Request Object
type ShowOptimizationTaskResultRequest struct {

	// 分子优化任务ID
	TaskId string `json:"task_id"`
}

func (o ShowOptimizationTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOptimizationTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowOptimizationTaskResultRequest", string(data)}, " ")
}
