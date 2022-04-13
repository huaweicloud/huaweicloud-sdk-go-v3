package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InstallMultiTasksRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *MultiTaskInitBody `json:"body,omitempty"`
}

func (o InstallMultiTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallMultiTasksRequest struct{}"
	}

	return strings.Join([]string{"InstallMultiTasksRequest", string(data)}, " ")
}
