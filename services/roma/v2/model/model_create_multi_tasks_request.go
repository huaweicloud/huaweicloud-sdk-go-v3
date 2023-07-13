package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMultiTasksRequest Request Object
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
