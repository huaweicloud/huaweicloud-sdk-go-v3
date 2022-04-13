package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreJobInstanceRequest struct {
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
