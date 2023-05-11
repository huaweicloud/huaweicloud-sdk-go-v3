package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobInstanceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	// 作业实例id.
	InstanceId string `json:"instance_id"`
}

func (o ShowJobInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceRequest", string(data)}, " ")
}
