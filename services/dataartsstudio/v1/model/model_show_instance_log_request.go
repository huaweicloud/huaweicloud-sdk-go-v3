package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceLogRequest Request Object
type ShowInstanceLogRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 例ID
	TaskId string `json:"task_id"`

	// 事例ID
	InstanceId string `json:"instance_id"`

	Body *JobLogRequest `json:"body,omitempty"`
}

func (o ShowInstanceLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceLogRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceLogRequest", string(data)}, " ")
}
