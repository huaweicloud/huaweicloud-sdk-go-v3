package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InstallMultiTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *MultiTaskInitBody `json:"body,omitempty" xml:"body"`
}

func (o InstallMultiTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallMultiTasksRequest struct{}"
	}

	return strings.Join([]string{"InstallMultiTasksRequest", string(data)}, " ")
}
