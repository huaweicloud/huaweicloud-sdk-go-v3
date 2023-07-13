package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreJobInstanceRequest Request Object
type RestoreJobInstanceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	// 作业实例id.
	InstanceId string `json:"instance_id"`
}

func (o RestoreJobInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreJobInstanceRequest", string(data)}, " ")
}
