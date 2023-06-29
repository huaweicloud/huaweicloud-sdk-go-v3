package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskInfoRequest Request Object
type UpdateTaskInfoRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 任务id
	TaskId string `json:"task_id"`

	Body *MetadataCollectionTask `json:"body,omitempty"`
}

func (o UpdateTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskInfoRequest", string(data)}, " ")
}
