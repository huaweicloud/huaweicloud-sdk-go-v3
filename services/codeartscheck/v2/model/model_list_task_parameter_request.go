package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskParameterRequest Request Object
type ListTaskParameterRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *ConfigTaskParameterBody `json:"body,omitempty"`
}

func (o ListTaskParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskParameterRequest struct{}"
	}

	return strings.Join([]string{"ListTaskParameterRequest", string(data)}, " ")
}
