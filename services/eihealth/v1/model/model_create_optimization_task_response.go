package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateOptimizationTaskResponse struct {

	// 分子优化任务ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOptimizationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOptimizationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateOptimizationTaskResponse", string(data)}, " ")
}
