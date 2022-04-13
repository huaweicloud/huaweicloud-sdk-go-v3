package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMultiTasksRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *MultiTaskRequestBody `json:"body,omitempty"`
}

func (o CreateMultiTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateMultiTasksRequest", string(data)}, " ")
}
