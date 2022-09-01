package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteMultiTaskMappingRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	// 组合任务映射唯一标识
	MappingId string `json:"mapping_id" xml:"mapping_id"`
}

func (o DeleteMultiTaskMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMultiTaskMappingRequest struct{}"
	}

	return strings.Join([]string{"DeleteMultiTaskMappingRequest", string(data)}, " ")
}
